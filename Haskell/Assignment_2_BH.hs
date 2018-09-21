-- CS300-SP17 Assignment 2: Barnes Hut Simulation
-- Deadline: 24 Feb 9pm
-- Submission: via LMS only

-- Update body working fine. Problem in either MakeQT or Tick function which does not return correct output

import System.Environment
import Data.List
import Graphics.Rendering.OpenGL hiding (($=))
import Graphics.UI.GLUT
import Control.Applicative
import Data.IORef
import Debug.Trace
import Data.Function (on)
--
-- PART 1: You are given many input files in the inputs directory and given 
-- code can read and parse the input files. You need to run "cabal update 
-- && cabal install glut" on your system to work on this assignment. Then
-- run "./ghc BH.hs && ./BH 1 < inputs/planets.txt" to run 1 iteration of
-- the algorithm and print updated positions. Replace 1 with any number to
-- run more iteration. You may also run it without an argument and it will
-- display the simulation using OpenGL in a window.
--
-- In the first part, you are to write the updateBody function to correctly 
-- update body after 1 unit of time has passed. You have to correctly 
-- update the position of a body by calculating the effect of all other
-- bodies on it. The placeholder implementation just moves the body left
-- without doing any physics. Read http://www.cs.princeton.edu/courses/
-- archive/fall03/cs126/assignments/nbody.html for help with physics. Try
-- simplyfying equations on paper before implementing them. You can compare
-- answers with the given binary solution.
--
-- Make helper functions as needed
constG :: Double
constG = 6.67*10**(-11)

radiusFUN :: Vec2 -> Vec2 -> Double -- planet1Pos -> planetsanet2 Pos -> ret Radius
radiusFUN (x1,y1) (x2,y2) = sqrt ( (x2-x1)^2 + (y2 - y1)^2 )


forceFUN :: Double -> Double -> Double -> Double -- MASS1 -> MASS2 -> Radius -> ret Force
forceFUN m1 m2 r 
    | r==0 = 0
    | otherwise = (constG * m1 * m2) / (r*r)
-- = (constG * m1 * m2) / (r*r)

force2FUN :: Double -> Vec2 -> Double -> Vec2 -- FORCE -> delXdelYTuple -> Radius -> Return ForceTuple
force2FUN justForce (delX, delY) r
    | r==0 = (0,0)
    | otherwise = (justForce * delX / r, justForce * delY / r) 

accelFUN :: Vec2 -> Double -> Vec2 -- fX,fY tuple -> mass -> ret accel
accelFUN (fX,fY) mass = (fX/mass, fY/mass)

velFUN :: Vec2 -> Double -> Vec2 -> Vec2 -- velocityTyple -> Time -> accelTuple -> ret velTuple
velFUN (vX,vY) delTime (aX,aY) = (vX + delTime*aX, vY + delTime*aY)

posFUN :: Vec2 -> Double -> Vec2 -> Vec2 -- positionTuple -> delTime -> VelTuple -> ret Position
posFUN (pX, pY) delTime (vX, vY) = (pX + delTime*vX, pY + delTime*vY)

type Vec2 = (Double, Double)
data Body = Body Vec2 Vec2 Double (Color3 Double)

updateBody :: (Foldable f) => f Body -> Body -> Body
updateBody mylovelylist (Body (posx,posy) vel mass clr) = 
    let (finalFx,finalFy) = 
                            foldr (\(Body (posx2,posy2) vel2 mass2 clr2) (iniX, iniY) -> 
                                let 
                                    radius = radiusFUN (posx,posy) (posx2,posy2)
                                    justForce = forceFUN mass mass2 radius
                                    (forceX, forceY) = force2FUN justForce (posx2-posx,posy2-posy) radius
                                in (iniX + forceX, iniY + forceY)
                            ) 
                            (0,0) mylovelylist -- remaining fold arguments  
    in 
    (
        let 
            accelTuple = accelFUN (finalFx,finalFy) mass
            velTuple = velFUN (vel) (1) (accelTuple)
            (finalPX, finalPY) = posFUN (posx,posy) (1) velTuple
        in Body (finalPX,finalPY) velTuple mass clr
    )
-- foldr (\x acc -> f x : acc) [] xs
-- updateBody _ (Body (posx,posy) vel mass clr) = Body (posx-100000000,posy) vel mass clr



-- PART 2: We wlil make a Quadtree to represent our universe. See 
-- http://www.cs.princeton.edu/courses/archive/fall03/cs126/assignmnefats/
-- barnes-hut.html for help on this tree. The QT structure has the the
-- length of one side of quadrant) for internal nodes. The makeQT function
-- has arguments: center, length of quadrant side, a function to find
-- coordinates of an element of the tree (e.g. extract position from a Body
-- object), a function to summarize a list of nodes (e.g. to calculate a
-- Body with position center of gravity and total mass), and the list of
-- nodes to put in tree.
--
-- Note that inserting all nodes at once is much easier than inserting one
-- by one. Think of creating the root node (given all nodes), and then
-- divide the nodes into quadrants and let recursive calls create the
-- appropriate trees. In this part you do not have to give a correct
-- implementation of the summary function.
--
-- Now make QT member of Foldable typeclass. See what function you are
-- required to implement. Once this is done, change the tick function below 
-- to create the tree and then pass it to updateBody function instead of a 
-- list of bodies. No change to updateBody function should be needed since
-- it works with any Foldable.

bodyPosFUN :: Body -> Vec2
bodyPosFUN (Body (posx,posy) vel mass clr) = (posx,posy)

summarizeFUN :: [Body] -> Body
summarizeFUN mylovelylist = 
    let net_mass = foldr (\(Body (posx,posy) vel mass clr) acc -> mass+acc) 0 mylovelylist
        (net_posx, net_posy) = foldr (\(Body (posx,posy) vel mass clr) (nX,nY) -> (posx*mass + nX, posy*mass + nY) ) (0,0) mylovelylist
    in Body (net_posx/net_mass,net_posy/net_mass) (0,0) net_mass (Color3 1 1 1)

getQuadrant :: Vec2 -> Vec2 -> Int
getQuadrant (cenx,ceny) (posx,posy)
    | posx <= cenx && posy > ceny = 1
    | posx > cenx && posy > ceny = 2
    | posx <= cenx && posy <= ceny = 3
    | posx > cenx && posy <= ceny = 4
    | otherwise = 1

getNewCenter :: Vec2 -> Double -> Int -> Vec2
getNewCenter (cenx, ceny) r quadNo
    | quadNo == 1 = (cenx - r/2, ceny + r/2)
    | quadNo == 2 = (cenx + r/2, ceny + r/2)
    | quadNo == 3 = (cenx - r/2, ceny - r/2)
    | quadNo == 4 = (cenx + r/2, ceny - r/2)

data QT a = Internal Double a (QT a,QT a,QT a,QT a) | Leaf a | Nil deriving (Show)

instance Foldable QT where
    foldr fun acc Nil = acc
    foldr fun acc (Leaf a) = fun a acc
    foldr fun acc (Internal doug alpha (beta, charlie, delta, echo)) = fun alpha (foldr fun (foldr fun (foldr fun ( foldr fun acc echo ) delta ) charlie ) beta) 

-- foldr fun (foldr fun a (foldr fun b (foldr fun c (foldr fun d (foldr fun acc e)))) 
-- Now make QT member of Foldable typeclass. See what function you are required to implement. Once this is done, change the tick function below 
-- to create the tree and then pass it to updateBody function instead of a list of bodies.

{-
taken from: http://learnyouahaskell.com/functors-applicative-functors-and-monoids

data Tree a = Empty | Leaf a | Node (Tree a) a (Tree a)

instance Foldable Tree where
   foldr fun acc Empty = acc
   foldr fun acc (Leaf a) = fun a acc
   foldr fun acc (Node l k r) = foldr fun (fun k (foldr fun acc r)) l

class Monoid m where  
    mempty :: m  
    mappend :: m -> m -> m  
    mconcat :: [m] -> m  
    mconcat = foldr mappend mempty  

instance F.Foldable Tree where  
    foldMap f Empty = mempty  
    foldMap f (Node x l r) = F.foldMap f l `mappend`  
                             f x           `mappend`  
                             F.foldMap f r 



instance Foldable Tree where
   foldMap f Empty = mempty
   foldMap f (Leaf x) = f x
   foldMap f (Node l k r) = foldMap f l `mappend` f k `mappend` foldMap f r

instance Foldable Tree where
   foldr f z Empty = z
   foldr f z (Leaf x) = f x z
   foldr f z (Node l k r) = foldr f (f k (foldr f z r)) l

-}


{-
randBody1 :: Body
randBody1 = Body (10,10) (0,0) 40 (Color3 0 0 1)
randBody2 :: Body
randBody2 = Body (20,20) (10,5) 80 (Color3 0 1 0)
randBody3 :: Body
randBody3 = Body (30,30) (0,15) 95 (Color3 0 1 1)
randBody4 :: Body
randBody4 = Body (-40,-40) (0,15) 95 (Color3 0 1 1)

testTree = Internal 5 randBody1 (Leaf randBody2, Nil, Nil, Leaf randBody3)
-}
-- Read the given link and see how center of quadtree nodes are set
-- You can add “deriving (Show)” at the end of “data QT...” line and then if “tree” is the name of the tree, you can prepend “trace (show tree)” to the “in” 
-- expression i.e. ““let tree = makeQT... in trace (show tree) ...” to see the tree and compare with the one in resultN files.
--                                              (e.g. extract position     (e.g. to calculate a Body with position 
--                                              from a Body object)        center of gravity and total mass)
-- ARGS => center -> length of quadrant side -> function to find coords -> function to summarize list of nodes -> list of nodes to put in tree -> return Quad Tree
--                                              of an element of tree
makeQT ::  Vec2 -> Double -> (Body->Vec2) -> ([Body]->Body) -> [Body] -> (QT Body)
makeQT _ _ _ _ [] = Nil
makeQT _ _ _ _ [x] = Leaf x
makeQT center@(cenx,ceny) radius bodyPosFUN summarizeFUN bodies =  
    let rootNode = summarizeFUN bodies
        list_one_bodies = [ (Body pos vel mass clr) | (Body pos vel mass clr) <- bodies, (getQuadrant center pos) == 1 ]
        list_two_bodies = [ (Body pos vel mass clr) | (Body pos vel mass clr) <- bodies, (getQuadrant center pos) == 2 ]
        list_three_bodies = [ (Body pos vel mass clr) | (Body pos vel mass clr) <- bodies, (getQuadrant center pos) == 3 ]
        list_four_bodies = [ (Body pos vel mass clr) | (Body pos vel mass clr) <- bodies, (getQuadrant center pos) == 4 ]
        newRad = radius/2
        cen1 = (cenx - newRad, ceny + newRad)
        cen2 = (cenx + newRad, ceny + newRad)
        cen3 = (cenx - newRad, ceny - newRad)   
        cen4 = (cenx + newRad, ceny - newRad)
    in Internal radius rootNode (  makeQT cen1 newRad bodyPosFUN summarizeFUN list_one_bodies,
                                   makeQT cen2 newRad bodyPosFUN summarizeFUN list_two_bodies,
                                   makeQT cen3 newRad bodyPosFUN summarizeFUN list_three_bodies,
                                   makeQT cen4 newRad bodyPosFUN summarizeFUN list_four_bodies ) 


        -- Body (posx,posy) vel mass clr

-- Once this is done, change the tick function below 
-- to create the tree and then pass it to updateBody function instead of a 
-- list of bodies.

-- This functions takes a set of bodies and returns an updated set of 
-- bodies after 1 unit of time has passed (dt=1)


tick :: Double -> [Body] -> [Body]
tick radius bodies = 
    let (Body (posx,posy) vel mass clr) = summarizeFUN bodies 
        center = (posx,posy)
    in fmap (updateBody (makeQT (posx,posy) radius bodyPosFUN summarizeFUN bodies)) bodies

-- in fmap (updateBody (makeQT (0.0,0.0) radius bodyPosFUN summarizeFUN bodies)) bodies
-- tick radius bodies = fmap (updateBody bodies) bodies

-- tick radius bodies =
--     let listBodies = [ (Body pos vel mass clr) | (Body pos vel mass clr) <- bodies]
--     in listBodies

-- PART 3: Now we create another datatype that contains a quadtree and a 
-- function which given radius and a summarized body (containing center of
-- gravity and total mass) returns true if the summarized body is a good
-- enough approximation. Use 0.5 as threshold.
--
-- Make a correct summarize function to pass to makeQT above and then make
-- BH an instance of Foldable typeclass as well. However this instance
-- should use the internal node if the predicate function returns true and
-- recurse only if it returns false. Make sure to recurse over a BH type
-- variable. If your implementation is correct, you will be as fast as the
-- provided binary BH2 on large inputs like galaxy1.txt
data BH a = BH (Double -> a -> Bool) (QT a)

---------------------------------------------------------------------------
-- You don't need to study the code below to work on the assignment
---------------------------------------------------------------------------
main :: IO ()
main = do
    (_,args) <- getArgsAndInitialize
    stdin <- getContents
    uncurry (mainChoice args) (parseInput stdin)

mainChoice :: [String] -> Double -> [Body] -> IO ()
mainChoice (iter:_) r bodies = putStr $ applyNtimes r bodies (read iter)
mainChoice [] r bodies = do
    createWindow "Barnes Hut"
    windowSize $= Size 700 700
    bodiesRef <- newIORef bodies
    ortho2D (-r) r (-r) r
    displayCallback $= (display r bodiesRef)
    addTimerCallback 10 (timer r bodiesRef)
    mainLoop

applyNtimes :: Double -> [Body] -> Int -> String
applyNtimes r bodies n = (unlines.map show) (iterate (tick r) bodies !! n)
 
parseInput :: String -> (Double, [Body])
parseInput input = 
    let (cnt:r:bodies) = lines input
    in (read r, map read (take (read cnt) bodies))

dispBody :: Body -> IO ()
dispBody (Body (x,y) _ _ rgb) = color rgb >> vertex (Vertex2 x y)

display :: Double -> IORef [Body] -> IO ()
display r bodiesRef = do
    clear [ColorBuffer]
    bodies <- get bodiesRef
    renderPrimitive Points (mapM_ dispBody bodies)
    flush

timer :: Double -> IORef [Body] -> IO ()
timer r bodiesRef = do
    postRedisplay Nothing
    bodies <- get bodiesRef
    bodiesRef $= tick r bodies 
    addTimerCallback 10 (timer r bodiesRef)

instance Read Body where
    readsPrec _ input = 
        let (x:y:vx:vy:m:r:g:b:rest) = words input
        in (\str -> [(Body (read x,read y) (read vx,read vy) (read m) 
            (Color3 ((read r)/255) ((read g)/255) ((read b)/255)), 
            unwords rest)]) input

instance Show Body where
    show (Body (x,y) (vx,vy) _ _) =
        "x=" ++ show x ++ " y=" ++ show y ++ " vx=" ++ 
            show vx ++ " vy=" ++ show vy


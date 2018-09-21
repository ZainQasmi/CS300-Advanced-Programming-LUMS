-- ---------------------------------------------------------------------
-- DNA Analysis
-- CS300 Spring 2017 Exam 1
--
-- >>> YOU ARE NOT ALLOWED TO IMPORT ANY LIBRARY
-- Functions available without import are okay
-- Making new helper functions is okay
--
-- >>> PARTS MUST BE DONE IN ORDER
-- Once "Check.exe" passes and you have not hard-coded for the test case, 
-- you are allowed to proceed to the next part but it does not guarantee 
-- full grade for the part. For grading, more test cases will be added.
--
-- >>> PARTS DO NOT HAVE EQUAL WEIGHT but earlier parts have more weight 
-- so start attempting without panicking about what's left
--
-- ---------------------------------------------------------------------
--
-- DNA can be thought of as a sequence of nucleotides. Each nucleotide is 
-- adenine, cytosine, guanine, or thymine. These are abbreviated as A, C, 
-- G, and T.
--
type DNA = [Char]

{-checkDNA :: [Char] -> Int
checkDNA [] = 0
checkDNA testA = length ([ oneChar | oneChar <- testA, oneChar == 'A' || oneChar == 'C' || oneChar == 'G' || oneChar == 'T'])
-}

{-
    | ('A' `elem` x || 'C' `elem` x  || 'G' `elem` x  || 'T' `elem` x  ) = checkDNA(xs)+1 
    | otherwise = checkDNA+1

-}
-- checkDNA (xs:xss) = [ "A" `elem`   x   |  x <- xs]

-- PART 1
-- Write a function to see if a given DNA is valid. A valid DNA consists 
-- of only A, C, G, and T and has more than one nucelotide
{-
isValid :: DNA -> Bool
isValid myDNA =
      let total_len = foldr (\oneStr acc -> acc + length(oneStr)) 0 myDNA
          len_filter = foldr (\oneStr acc2 -> acc2 + length(checkDNA oneStr)) 0 myDNA
      in if total_len == len_filter then True else False
-}


{-
isValid :: DNA -> Bool
isValid myDNA =

      let total_len = length (myDNA)
          len_filter = checkDNA myDNA
      in if total_len == len_filter then True else False

-}
-- testA :: DNA
-- testA = "AACCTTGG"
-- testA = [["AACCTTGG"],["ACTGCATG"], ["ACTACACC"], ["ATATTATA"]]

checkDNA :: [Char] -> Int
checkDNA [] = 0
checkDNA (x:xs)
    | x == 'A' || x == 'C' || x == 'G' || x == 'T' =  checkDNA(xs)+1 
    | otherwise = checkDNA xs

checkDNA2 :: [Char] -> Int
checkDNA2 [] = 0
checkDNA2 (x:xs)
    | x == 'C' || x == 'G' =  checkDNA2(xs)+1 
    | otherwise = checkDNA2 xs

isValid :: DNA -> Bool
isValid myDNA = if (length myDNA > 1) then (checkDNA myDNA) == length myDNA else False


-- PART 2
-- You have to compute the GC content of DNA data. The GC content of DNA 
-- is the percentage of nucleotides that are either G or C. It can be used 
-- in determining classification of species. The answer can be from 0 to 100
{-
gcContent :: DNA -> Int
gcContent myDNA =
        let tot = checkDNA myDNA
            chkGC = checkDNA2 myDNA
            frac = tot `div` chkGC
            frac2 = frac*100
            frac3 = frac2 `div` chkGC
        in (frac3)gcContent :: DNA -> Int
        -}
gcContent :: DNA -> Int 
gcContent myDNA = ((checkDNA2 myDNA)*100) `div` length myDNA

-- PART 3
-- We want to also compute the AT content. The AT content is the percentage 
-- of nucleotides that are A or T. However, instead of writing another 
-- specific function, we want a generic function that given a list of 
-- nucleotides, finds their percentage in the DNA. The answer can be from 0 
-- to 100.

content :: [Char] -> DNA -> Int
content listNuc myDNA =  
    let a = [ oneStr | oneStr <- listNuc, oneLet <- myDNA, oneStr == oneLet ]
        b = length a
    in (b*100) `div` (length myDNA)
-- PART 4
-- Another way to classify DNA is using the AT/GC ratio i.e. the ratio of 
-- nucleotides that are A or T to those that are G or C. Give the answer as 
-- a percentage that will be a non-negative number. All AT should return Nothing.
atgcRatio :: DNA -> Maybe Int
atgcRatio myDNA =
    let gc = gcContent myDNA
        at = 100 - gcContent myDNA
    in if at == 100 then Nothing else Just ((at `div` gc )*100)

-- PART 5 
-- Write a function to find the maximum from a list of elements on which 
-- you may call "max" to find the bigger of two elements. The first argument
-- is the value to be used if the list is empty.
maxList :: (Ord a) => a -> [a] -> a
maxList ifEmptyArg [] = ifEmptyArg
maxList ifEmptyArg [oneVal] = oneVal
maxList ifEmptyArg myList = foldr (\oneInt acc2 -> (max oneInt acc2)) (myList!!0) myList


-- PART 6
-- Now we want to calculate how alike are two DNA strands. We will try to 
-- align the DNA strands. An aligned nucleotide gets 3 points, a misaligned
-- gets 2 points, and inserting a gap in one of the strands gets 1 point. 
-- Since we are not sure how the next two characters should be aligned, we
-- try three approaches and pick the one that gets us the maximum score.
-- 1) Align or misalign the next nucleotide from both strands
-- 2) Align the next nucleotide from first strand with a gap in the second
-- 3) Align the next nucleotide from second strand with a gap in the first
-- In all three cases, we calculate score of leftover strands from recursive
-- call and add the appropriate penalty.
drop_nth::String->Int->String
drop_nth s n = f1 0 s
        where f1 c str
                | c == (length str) = [' ']
                | (mod c n) /= 0 = [str !! c] ++ f1 (c+1) str
                | otherwise = f1 (c+1) str 

rem_kth::(Eq a)=> [a]->Int->[a]
rem_kth [] _ = []
rem_kth (x:xs) n
        | n /= 0 = x:rem_kth xs (n-1)
        | otherwise = rem_kth xs (n-1)

num_oc::(Eq a)=> a -> [a] -> Int
num_oc _ [] = 0
num_oc x a = foldr (\_ n-> 1+n) 0 (filter (\p-> x==p) a)

scoreHelp :: [Char] -> [Char] -> Int
scoreHelp [] [] = 3
scoreHelp [] (y:ys) = 1 + scoreHelp [] ys
scoreHelp (x:xs) [] = 1 + scoreHelp xs []
scoreHelp (x:xs) (y:ys) = if x==y then (3 + scoreHelp xs ys) else (2 + scoreHelp xs ys)




score :: DNA -> DNA -> Int
score myDNA1 myDNA2 = 
    let str1 = if length(myDNA1) >= length(myDNA2) then myDNA1 else myDNA2
        str2 = if length(myDNA1) < length(myDNA2) then myDNA1 else myDNA2
        aggregate = scoreHelp str1 str2
    in aggregate

{-
score :: DNA -> DNA -> Int
score = undefined
-}
-- PART 7
-- Write a function that takes a list of DNA strands and returns a list
-- of nodes. For each DNA strand, make a separate node with zero score 
-- in the Int field and Nil sub-trees

data DNATree = Node Int DNA DNATree DNATree | Nil deriving (Ord, Show, Eq)

makeDNATree :: [DNA] -> [DNATree]
makeDNATree [] = []
--                           Node Int DNA DNATree DNATree
makeDNATree myTree@(x:xs) = [Node 0   x   Nil     Nil] ++ makeDNATree xs


-- PART 8
-- Write a function to find the two closest DNA strands and make a new 
-- Node containing their score and the smaller DNA (using min). One 
-- approach is to make a new Node for every pair of nodes and than use
-- maximum to find the one with the highest score. Your maximum function
-- will work on Nodes. 

-- maxSimilar = undefined
-- maxSimilar (x1:x2:xs) = 

{-
split :: [a] -> ([a],[a])
split [] = ([],[])
split [x] = ([x],[])
split (x1:x2:xys) = 
    let (y,z) = split xys
    in ((x1:y), (x2:z))
-}

testA = ["AACCTTGG","ACTGCATG", "ACTACACC", "ATATTATA"]

closestDNA :: [DNA] -> [Int]
closestDNA [] = [0]
closestDNA [oneVal] = [0]
closestDNA (x1:x2:xs) = (score x1 x2):closestDNA (x2:xs)


{-getTuple :: Int -> [DNA] -> (DNA,DNA)
getTuple maxScore myList = 
        let (x1,x2) = 

-}
findMatch :: [Int] -> Int -> Int
findMatch maxScoreList toChk= 
    let maxScorePos =  foldr (\oneInt acc -> if oneInt == maximum (maxScoreList) then acc-1 else acc+1) 0 maxScoreList
    in maxScorePos
{-
findIndex :: [Int] -> Int -> Int
findIndex [] _ = 0
findIndex (x:xs) toChk
    | x == toChk = (findIndex (xs) (toChk))
    | x /= toChk = (findIndex (xs) (toChk)) +1
-}

findIndex :: [Int] -> Int -> Int
findIndex [] _ = 0
findIndex (x:xs) toChk
    | x == toChk = (findIndex (xs) (toChk))
    | x /= toChk = (findIndex (xs) (toChk)) +1



maxSimilar :: [DNATree] -> DNATree
maxSimilar = undefined

{-
############ Closest DNA finds max score 
-}

{-
maxSimilar :: [DNATree] -> DNATree
maxSimilar treeList = 
    let listScores = [ score | (Node score x left right) <- treeList]
        listDNA = [ x | (Node score x left right) <- treeList]


        -- a2 = [Node 0 x2 Nil Nil]
    in head treeList




-}
-- PART 9
-- Write a function that take a list of DNA strands and replaces the two
-- closest ones with one Node made from maxSimilar. The returned list will
-- have one less element than the input list.
evolStep :: [DNATree] -> [DNATree]
evolStep = undefined

-- PART 10
-- Repeat the above function until you are left with a single Node 
-- representing all DNA strands you started with.
makeEvolTree :: [DNATree] -> DNATree
makeEvolTree = undefined

-- PART 11
-- Lets try to neatly print the DNATree. Each internal node should show the 
-- match score while leaves should show the DNA strand. In case the DNA strand 
-- is more than 10 characters, show only the first seven followed by "..." 
-- The tree should show like this for an evolution tree of
-- ["AACCTTGG","ACTGCATG", "ACTACACC", "ATATTATA"]
--
-- 20
-- +---ATATTATA
-- +---21
--     +---21
--     |   +---ACTGCATG
--     |   +---ACTACACC
--     +---AACCTTGG
--
-- Make helper functions as needed. It is a bit tricky to get it right. One
-- hint is to pass two extra string, one showing what to prepend to next 
-- level e.g. "+---" and another to prepend to level further deep e.g. "|   "
draw :: DNATree -> String
draw = undefined

-- PART 12
-- Our score function is inefficient due to repeated calls for the same 
-- suffixes. Lets make a dictionary to remember previous results. First you
-- will consider the dictionary as a list of tuples and write a lookup
-- function. Return Nothing if the element is not found. Also write the 
-- insert function. You can assume that the key is not already there.
type Dict a b = [(a,b)]

lookupDict :: (Eq a) => a -> Dict a b -> Maybe b
lookupDict = undefined

insertDict :: (Eq a) => a -> b -> (Dict a b)-> (Dict a b)
insertDict = undefined

-- PART 13
-- We will improve the score function to also return the alignment along
-- with the score. The aligned DNA strands will have gaps inserted. You
-- can represent a gap with "-". You will need multiple let expressions 
-- to destructure the tuples returned by recursive calls.
alignment :: String -> String -> (Int,String,String)
alignment = undefined

-- PART 14
-- We will now pass a dictionary to remember previously calculated scores 
-- and return the updated dictionary along with the result. Use let 
-- expressions like the last part and pass the dictionary from each call
-- to the next. Also write logic to skip the entire calculation if the 
-- score is found in the dictionary. You need just one call to insert.
type ScoreDict = Dict (DNA,DNA) Int

scoreMemo :: (DNA,DNA) -> ScoreDict -> (ScoreDict,Int)
scoreMemo = undefined

-- PART 15
-- In this part, we will use an alternate representation for the 
-- dictionary and rewrite the scoreMemo function using this new format.
-- The dictionary will be just the lookup function so the dictionary 
-- can be invoked as a function to lookup an element. To insert an
-- element you return a new function that checks for the inserted
-- element and returns the old dictionary otherwise. You will have to
-- think a bit on how this will work. An empty dictionary in this 
-- format is (\_->Nothing)
type Dict2 a b = a->Maybe b
insertDict2 :: (Eq a) => a -> b -> (Dict2 a b)-> (Dict2 a b)
insertDict2 = undefined

type ScoreDict2 = Dict2 (DNA,DNA) Int

scoreMemo2 :: (DNA,DNA) -> ScoreDict2 -> (ScoreDict2,Int)
scoreMemo2 = undefined


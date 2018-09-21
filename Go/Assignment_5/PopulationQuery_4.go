package main

import (
    "fmt"
    "os"
    "strconv"
    "math"
    "encoding/csv"
    "sync"
)

type CensusGroup struct {
    population int
    latitude, longitude float64
}

func ParseCensusData(fname string) ([]CensusGroup, error) {
    file, err := os.Open(fname)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    records, err := csv.NewReader(file).ReadAll()
    if err != nil {
        return nil, err
    }
    censusData := make([]CensusGroup, 0, len(records))

    for _, rec := range records {
        if len(rec) == 7 {
            population, err1 := strconv.Atoi(rec[4])
            latitude, err2 := strconv.ParseFloat(rec[5], 64)
            longitude, err3 := strconv.ParseFloat(rec[6], 64)
            if err1 == nil && err2 == nil && err3 == nil {
                latpi := latitude * math.Pi / 180
                latitude = math.Log(math.Tan(latpi) + 1 / math.Cos(latpi))
                censusData = append(censusData, CensusGroup{population, latitude, longitude})
            }
        }
    }

    return censusData, nil
}

var popTotalG int
var maxLatitudeG float64 
var minLatidudeG float64
var maxLongitudeG float64
var minLongitudeG float64


func findLimits(data []CensusGroup, mu *sync.Mutex) {
    if len(data) <= 10000 {
        for i := 0; i<len(data);i++ {
            if data[i].latitude > maxLatitudeG {
                maxLatitudeG = data[i].latitude
            } 
            if data[i].latitude < minLatidudeG {
                minLatidudeG = data[i].latitude
            }
            if data[i].longitude > maxLongitudeG {
                maxLongitudeG = data[i].longitude
            } 
            if data[i].longitude < minLongitudeG {
                minLongitudeG = data[i].longitude
            }
            mu.Lock()
            popTotalG += data[i].population
            mu.Unlock()
        }
        return
    }
    mid := len(data)/2
    done := make(chan bool)
    go func() {
        findLimits(data[:mid], mu)
        done <- true
    }()
    findLimits(data[mid:], mu)
    <-done
    return
}

func findParallel(data []CensusGroup, queryWest float64, querySouth float64, queryEast float64, queryNorth float64, population *int, mu *sync.Mutex) {
    if len(data) <= 10000 {
        for i := 0; i<len(data);i++ {
            if data[i].longitude >= queryWest && data[i].longitude <= queryEast && data[i].latitude >= querySouth && data[i].latitude <= queryNorth {
                mu.Lock()
                *population += data[i].population
                mu.Unlock()
            }
        }
        return
    }
    mid := len(data)/2
    done := make(chan bool)
    go func() {
        findParallel(data[:mid], queryWest, querySouth, queryEast, queryNorth, population, mu)
        done <- true
    }()
    findParallel(data[mid:], queryWest, querySouth, queryEast, queryNorth, population, mu)
    <-done
    return
}

// func teller(deposits chan int, balances chan int) {
//     balance := 0
//     for {
//         select {
//         case amount := <-deposits:
//             balance = balance + amount
//         case balances<- balance:
//         }
//     }
// }

/* Helper function for v4 */
func arrayAccumulate(target [][]int, src [][]int) {
    for i := 0; i < len(target); i++ {
        for j := 0; j < len(target[0]); j++ {
            target[i][j] += src[i][j]
        }
    }
}


var lolGrid [][]int

// func smartParallel(xdim,ydim) {
//     for i := 0; i < ydim; i++ {
//         lolGrid[i] = make([]int, xdim)
//     }
//     // balance := 0
//     for {
//         select {
//         case amount := <-deposits:
//             balance = balance + amount
//         case balances<- balance:
//         }
//     }
// }

func main () {
    if len(os.Args) < 4 {
        fmt.Printf("Usage:\nArg 1: file name for input data\nArg 2: number of x-dim buckets\nArg 3: number of y-dim buckets\nArg 4: -v1, -v2, -v3, -v4, -v5, or -v6\n")
        return
    }
    fname, ver := os.Args[1], os.Args[4]
    xdim, err := strconv.Atoi(os.Args[2])
    if err != nil {
        fmt.Println(err)
        return
    }
    ydim, err := strconv.Atoi(os.Args[3])
    if err != nil {
        fmt.Println(err)
        return
    }
    censusData, err := ParseCensusData(fname)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("%v%v\n", "censusData :: ", len(censusData));

    grid2D := make([][]int, xdim)
    for i := 0; i < xdim; i++ {
        grid2D[i] = make ([]int, ydim)
    }

    var queryWest float64
    var queryEast float64
    var querySouth float64
    var queryNorth float64

    var x_chunk float64
    var y_chunk float64

    // Some parts may need no setup code
    switch ver {
    case "-v1":
        // YOUR SETUP CODE FOR PART 1

    case "-v2":
        // YOUR SETUP CODE FOR PART 2


    case "-v3":
        // YOUR SETUP CODE FOR PART 3
    case "-v4":
        // YOUR SETUP CODE FOR PART 4
        popTotalG = 0
        maxLatitudeG = censusData[0].latitude
        minLatidudeG = censusData[0].latitude
        maxLongitudeG = censusData[0].longitude
        minLongitudeG = censusData[0].longitude
        
        var mu sync.Mutex
        findLimits(censusData, &mu)

        fmt.Printf("%v%v\n", "xdim :: ",xdim);
        fmt.Printf("%v%v\n", "ydim :: ",ydim);
        fmt.Printf("%v%v\n", "popTotalG :: ", popTotalG);
        fmt.Printf("%v%v\n%v%v\n%v%v\n%v%v\n", "maxLatitudeG :: ",maxLatitudeG, "minLatitudeG :: ", minLatidudeG,"maxLongitudeG :: ",maxLongitudeG,"minLongitudeG :: ", minLongitudeG );

        var x_chunk float64 = ((-minLongitudeG) - (-maxLongitudeG)) / float64(xdim)
        var y_chunk float64 = (maxLatitudeG - minLatidudeG) / float64(ydim)


        // TASK 4 COMMENTED BELOW. DOES NOT WORK

        // done := make(chan bool)
        // deposits := make(chan [][]int)
        // outputChannel := make(chan [][]int)
        // smartParallel(grid2D, channels, xdim, ydim, x_chunk, y_chunk)

        // Version 3: Step 1
        for k := 0; k<len(censusData); k++ {
            for i := 0; i<xdim;i++ {
                for j := 0; j<ydim; j++ {

                    queryWest = minLongitudeG + float64(x_chunk)*float64(i)
                    queryEast = minLongitudeG + float64(x_chunk)*float64(i)+x_chunk
                    querySouth = minLatidudeG + float64(y_chunk)*float64(j)
                    queryNorth = minLatidudeG + float64(y_chunk)*float64(j)+y_chunk

                    if censusData[k].longitude >= queryWest && censusData[k].longitude <= queryEast && censusData[k].latitude >= querySouth && censusData[k].latitude <= queryNorth {
                        grid2D[i][j] = grid2D[i][j] + censusData[k] .population
                    }

                }
            }            
        }

        // Version 3: Step 2
        var tempi = 0
        var tempj = 0

        for j := ydim-1; j>=0; j-- {
            for i := 0; i<xdim; i++ {            
                tempi = i-1
                tempj = j+1
                if ((tempi < 0) && tempj<=ydim-1) {
                    grid2D[i][j] = grid2D[i][j] + grid2D[i][tempj]
                } else if (tempi>=0 && (tempj >= ydim)) {
                    grid2D[i][j] = grid2D[i][j] + grid2D[tempi][j]
                } else if ((tempi < 0) && (tempj >= ydim)) {
                    grid2D[i][j] = grid2D[i][j]
                } else {
                    // fmt.Printf("%v\t%v\n", tempi, tempj)
                    grid2D[i][j] = grid2D[i][j] + grid2D[tempi][j] + grid2D[i][tempj] - grid2D[tempi][tempj]
                }
            }
        }

        fmt.Printf("%v%v\n", "pls be the right pop => grid2D[xdim-1][0] :: ",grid2D[xdim-1][0])


    case "-v5":
        // YOUR SETUP CODE FOR PART 5
    case "-v6":
        // YOUR SETUP CODE FOR PART 6
    default:
        fmt.Println("Invalid version argument")
        return
    }

    for {
        var west, south, east, north int
        n, err := fmt.Scanln(&west, &south, &east, &north)
        if n != 4 || err != nil || west<1 || west>xdim || south<1 || south>ydim || east<west || east>xdim || north<south || north>ydim {
            break
        }

        // west = west - 1     // ahh never mind
        // south = south - 1   // ahh never mind

        var population int //init with 0
        var percentage float64



        switch ver {
        case "-v1":
            // YOUR QUERY CODE FOR PART 1
        case "-v2":
            // YOUR QUERY CODE FOR PART 2
            var queryWest float64 = minLongitudeG + float64(x_chunk)*float64(west)
            var querySouth float64 = minLatidudeG + float64(y_chunk)*float64(south)
            var queryEast float64 = minLongitudeG + float64(x_chunk)*float64(east)
            var queryNorth float64 = minLatidudeG + float64(y_chunk)*float64(north)

            population = 0
            var mu2 sync.Mutex
            findParallel(censusData, queryWest, querySouth, queryEast, queryNorth, &population, &mu2)
            percentage = (float64(population)/float64(popTotalG)) * float64(100)
        case "-v3":
            // YOUR QUERY CODE FOR PART 3

        case "-v4":
            // YOUR QUERY CODE FOR PART 4
            var above, left, above_left int
            if north+1 <= ydim {
                above = grid2D[east-1][north]
            } else {
                above = 0
            }
            if west-1 >= 1 {
                left = grid2D[west-2][south-1]
            } else {
                left = 0
            }
            if west-1 >= 1 && north+1 <= ydim {
                above_left = grid2D[west-2][north]
            } else {
                above_left = 0
            }
            population = grid2D[east-1][south-1] - above - left + above_left
            percentage = (float64(population)/float64(popTotalG)) * float64(100)
        case "-v5":
            // YOUR QUERY CODE FOR PART 5
        case "-v6":
            // YOUR QUERY CODE FOR PART 6
        }
        fmt.Printf("%v %.2f%%\n", population, percentage)
    }
}

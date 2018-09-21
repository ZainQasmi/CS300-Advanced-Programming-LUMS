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


var populationTotal int
var maxLatitude float64 
var minLatidude float64
var maxLongitude float64
var minLongitude float64

func findLimits(data []CensusGroup, mu *sync.Mutex) {
    if len(data) <= 10000 {
        for i := 0; i<len(data);i++ {
            if data[i].latitude > maxLatitude {
                maxLatitude = data[i].latitude
            } 
            if data[i].latitude < minLatidude {
                minLatidude = data[i].latitude
            }
            if data[i].longitude > maxLongitude {
                maxLongitude = data[i].longitude
            } 
            if data[i].longitude < minLongitude {
                minLongitude = data[i].longitude
            }
            mu.Lock()
            populationTotal += data[i].population
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

func parallelStepOne (data []CensusGroup, grid2D *[][]int, lockArr *[][]sync.Mutex ,xdim int, ydim int, x_chunk float64, y_chunk float64) {

    if len(data) <= 10000 {
        for k := 0; k<len(data); k++ {
            // for i := xdim-1; i<=0;i-- {
            for i := 0; i<xdim;i++ {
                // for j := ydim-1; j<=0; j-- {
                for j := 0; j<ydim; j++ {
                    var queryWest, queryEast, querySouth, queryNorth float64

                    queryWest = minLongitude + float64(x_chunk)*float64(i)
                    queryEast = minLongitude + float64(x_chunk)*float64(i)+x_chunk
                    querySouth = minLatidude + float64(y_chunk)*float64(j)
                    queryNorth = minLatidude + float64(y_chunk)*float64(j)+y_chunk

                    if data[k].longitude >= queryWest && data[k].longitude <= queryEast && data[k].latitude >= querySouth && data[k].latitude <= queryNorth {
                        (*lockArr)[i][j].Lock()
                        (*grid2D)[i][j] = (*grid2D)[i][j] + data[k].population
                        (*lockArr)[i][j].Unlock()
                    }

                }
            }            
        }
        return
    }

    mid := len(data)/2
    done := make(chan bool)
    go func() {
        parallelStepOne (data[:mid], grid2D, lockArr, xdim, ydim, x_chunk, y_chunk)
        done <- true
    }()
    parallelStepOne (data[mid:], grid2D, lockArr, xdim, ydim, x_chunk, y_chunk)
    <-done
    return

}

// func parallelStepTwo (data *[][]int, lockArr *[][]sync.Mutex ,xdim int, ydim int) {
    
//     if len(*data) <= 2 {
//         for i := 0; i<xdim-1;i++ {            
//             for j := 0; j<ydim; j++ {
//                 // fmt.Printf("%v\t%v\n", i, j)
//                 (*data)[i+1][j] = (*data)[i][j] + (*data)[i+1][j] 
//             }
//         }
//         return
//     }

//     mid := len(*data)/2
//     done := make(chan bool)
//     go func() {
//         parallelStepTwo (data[:mid][:], lockArr, xdim, ydim)
//         done <- true
//     }()
//     parallelStepTwo (data[mid:][:], lockArr, xdim, ydim)
//     <-done
//     return

// }

// func parallelStepTwo (data *[]int, lockArr *[][]sync.Mutex ,xdim int, ydim int) {
    
//     if len(*data) <= 2 {
//         // for i := 0; i<xdim-1;i++ {            
//             for j := 0; j<ydim; j++ {
//                 // fmt.Printf("%v\t%v\n", i, j)
//                 // (*data)[i+1][j] = (*data)[i][j] + (*data)[i+1][j] 
//                 (*data)[j] = (*data)[j] + (*data)[j] 
//             }
//         // }
//         return
//     }

//     mid := len(*data)/2
//     done := make(chan bool)
//     go func() {
//         parallelStepTwo (data[:mid], lockArr, xdim, ydim)
//         done <- true
//     }()
//     parallelStepTwo (data[mid:], lockArr, xdim, ydim)
//     <-done
//     return

// }

func parallelStepTwo_v2(grid2D [][]int, temp int, ydim int, done chan bool){
    for i := ydim-2; i>=0; i-- {
        grid2D[temp][i] = grid2D[temp][i+1] + grid2D[temp][i]
    }
    done <- true
}


func parallelStepTwo_v1(grid2D [][]int, temp int, xdim int, done chan bool){
    for i := 1; i<xdim;i++ {
        grid2D[i][temp] = grid2D[i-1][temp] + grid2D[i][temp]
    }
    done <- true
}

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

    // Some parts may need no setup code
       
    grid2D := make([][]int, xdim)
    for i := 0; i < xdim; i++ {
        grid2D[i] = make ([]int, ydim)
    }

    Lock2D := make([][]sync.Mutex, xdim)
    for i := 0; i < xdim; i++ {
        Lock2D[i] = make ([]sync.Mutex, ydim)
    }

    Lock2D_2 := make([][]sync.Mutex, xdim)
    for i := 0; i < xdim; i++ {
        Lock2D_2[i] = make ([]sync.Mutex, ydim)
    }

    switch ver {
    case "-v1":
        // YOUR SETUP CODE FOR PART 1
    case "-v2":
        // YOUR SETUP CODE FOR PART 2
    case "-v3":
        // YOUR SETUP CODE FOR PART 3
    case "-v4":
        // YOUR SETUP CODE FOR PART 4   
    case "-v5":
        // YOUR SETUP CODE FOR PART 5

        var x_chunk float64
        var y_chunk float64

        populationTotal = 0
        maxLatitude = censusData[0].latitude
        minLatidude = censusData[0].latitude
        maxLongitude = censusData[0].longitude
        minLongitude = censusData[0].longitude

        var mu sync.Mutex
        findLimits(censusData, &mu)

        fmt.Printf("%v%v\n", "xdim :: ",xdim);
        fmt.Printf("%v%v\n", "ydim :: ",ydim);
        fmt.Printf("%v%v\n", "populationTotal :: ", populationTotal);
        fmt.Printf("%v%v\n%v%v\n%v%v\n%v%v\n", "maxLatitude :: ",maxLatitude, "minLatitude :: ", minLatidude,"maxLongitude :: ",maxLongitude,"minLongitude :: ", minLongitude );

        x_chunk = (maxLongitude - minLongitude) / float64(xdim)
        y_chunk = (maxLatitude - minLatidude) / float64(ydim)

        parallelStepOne (censusData, &grid2D, &Lock2D, xdim, ydim, x_chunk, y_chunk)

        // Version 3: Step 1
        // var queryWest float64
        // var queryEast float64
        // var querySouth float64
        // var queryNorth float64  
        // for k := 0; k<len(censusData); k++ {
        //     for i := 0; i<xdim;i++ {
        //         for j := 0; j<ydim; j++ {

        //             queryWest = minLongitude + float64(x_chunk)*float64(i)
        //             queryEast = minLongitude + float64(x_chunk)*float64(i)+x_chunk
        //             querySouth = minLatidude + float64(y_chunk)*float64(j)
        //             queryNorth = minLatidude + float64(y_chunk)*float64(j)+y_chunk

        //             if censusData[k].longitude >= queryWest && censusData[k].longitude <= queryEast && censusData[k].latitude >= querySouth && censusData[k].latitude <= queryNorth {
        //                 grid2D[i][j] = grid2D[i][j] + censusData[k] .population
        //             }

        //         }
        //     }            
        // }

        // Version 3: Step 2
        // var tempi = 0
        // var tempj = 0
        // for j := ydim-1; j>=0; j-- {
        //     for i := 0; i<xdim; i++ {            
        //         tempi = i-1
        //         tempj = j+1
        //         if ((tempi < 0) && tempj<=ydim-1) {
        //             grid2D[i][j] = grid2D[i][j] + grid2D[i][tempj]
        //         } else if (tempi>=0 && (tempj >= ydim)) {
        //             grid2D[i][j] = grid2D[i][j] + grid2D[tempi][j]
        //         } else if ((tempi < 0) && (tempj >= ydim)) {
        //             grid2D[i][j] = grid2D[i][j]
        //         } else {
        //             // fmt.Printf("%v\t%v\n", tempi, tempj)
        //             grid2D[i][j] = grid2D[i][j] + grid2D[tempi][j] + grid2D[i][tempj] - grid2D[tempi][tempj]
        //         }
        //     }
        // }

        doneA := make(chan bool)
        for i := ydim-1; i>=0 ; i-- {
          go parallelStepTwo_v1(grid2D, i, xdim, doneA)
        }

        for i := ydim-1; i>=0; i-- {
          <-doneA  
        }
        doneB := make(chan bool)
        for i := 0; i < xdim; i++ {
          go parallelStepTwo_v2(grid2D, i, ydim, doneB)
        }
        for i := 0; i < xdim; i++ {
          <-doneB
        }

        // done := make(chan bool)

        // for i := 0; i<xdim-1;i++ {
        //     go func (i int) {
        //         for j := 0; j<ydim; j++ {
        //             // fmt.Printf("%v\t%v\n", i, j)
        //             grid2D[i+1][j] = grid2D[i][j] + grid2D[i+1][j] 
        //         }
        //         done <- true
        //     } (i)
        //     <-done
        // }

        // for i := 0; i<xdim; i++ {
        //     go func (i int) {
        //         for j:=ydim-1; j>0; j-- {
        //             // fmt.Printf("%v\t%v\n", i, j)
        //             grid2D[i][j-1] = grid2D[i][j-1] + grid2D[i][j] 
        //         }
        //         done <- true
        //     } (i)
        //     <-done
        // }

        fmt.Printf("%v%v\n", "pls be the right pop => grid2D[xdim-1][0] :: ",grid2D[xdim-1][0])
        fmt.Printf("%v%v\n", "diff :: ",populationTotal - grid2D[xdim-1][0])

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

        var population int
        var percentage float64

        switch ver {
        case "-v1":
            // YOUR QUERY CODE FOR PART 1
        case "-v2":
            // YOUR QUERY CODE FOR PART 2
        case "-v3":
            // YOUR QUERY CODE FOR PART 3
        case "-v4":
            // YOUR QUERY CODE FOR PART 4
        case "-v5":
            // YOUR QUERY CODE FOR PART 5
            // if north-1 < 0 && west-1 >= 0 && south-1 >= 0{
            //     population = grid2D[east][south] - grid2D[west-1][south] + grid2D[west-1][south-1]
            // } else if west-1 < 0 && north-1 >= 0{
            //     population = grid2D[east][south] - grid2D[east][north-1]
            // } else if north-1 <0 && west-1 <0 && south-1 <0 {
            //     population = grid2D[east][south]
            // } else {
            //     population = grid2D[east][south] - grid2D[east][north-1] - grid2D[west-1][south] + grid2D[west-1][south-1]  
            // } 
            // percentage = (float64(population)/float64(populationTotal)) * float64(100)

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
            percentage = (float64(population)/float64(populationTotal)) * float64(100)

        case "-v6":
            // YOUR QUERY CODE FOR PART 6
        }

        fmt.Printf("%v %.2f%%\n", population, percentage)
    }
}

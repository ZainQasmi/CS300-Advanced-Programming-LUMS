package main

import (
    "fmt"
    "os"
    "strconv"
    "math"
    "encoding/csv"
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
    var populationTotal int = 0
    var maxLatitude float64 = censusData[0].latitude
    var minLatidude float64 = censusData[0].latitude
    var maxLongitude float64 = censusData[0].longitude
    var minLongitude float64 = censusData[0].longitude

    for i := 0; i < len(censusData); i++ { // Can start with i=1 to reduce computation
        if censusData[i].latitude > maxLatitude {
            maxLatitude = censusData[i].latitude
        } 
        if censusData[i].latitude < minLatidude {
            minLatidude = censusData[i].latitude
        }
        if censusData[i].longitude > maxLongitude {
            maxLongitude = censusData[i].longitude
        } 
        if censusData[i].longitude < minLongitude {
            minLongitude = censusData[i].longitude
        }
        populationTotal += censusData[i].population
    }

    fmt.Printf("%v%v\n", "xdim :: ",xdim);
    fmt.Printf("%v%v\n", "ydim :: ",ydim);
    fmt.Printf("%v%v\n", "populationTotal :: ", populationTotal);


    fmt.Printf("%v%v\n%v%v\n%v%v\n%v%v\n", "maxLatitude :: ",maxLatitude, "minLatitude :: ", minLatidude,"maxLongitude :: ",maxLongitude,"minLongitude :: ", minLongitude );

    // var x_chunk float64 = ((-minLongitude) - (-maxLongitude)) / float64(xdim)
    // var y_chunk float64 = (maxLatitude - minLatidude) / float64(ydim)

    var x_chunk float64 = (maxLongitude - minLongitude) / float64(xdim)
    var y_chunk float64 = (maxLatitude - minLatidude) / float64(ydim)

    // fmt.Printf("%v%v\n", "xchunk :: ",minLongitude + float64(x_chunk)*float64(xdim));
    // fmt.Printf("%v%v\n", "ychunk :: ",minLatidude + float64(y_chunk)*float64(ydim));
    // var grid [100][50][2]float64

    // Initialize 3D Array of size xdim x ydim x 2
    // Last dimension has longitude in [0] and latitude in [1]
    // grid := make([][][]float64, xdim)
    // for i := 0; i < xdim; i++ {
    //     grid[i] = make ([][]float64, ydim)
    //     for j:= 0; j < ydim; j++ {
    //         grid[i][j] = make([]float64, 2)
    //     }
    // }

    // Divide large rectangle into equal chunks of latitudes and longitudes
    // then store them in grid
    // for i := 0; i<xdim;i++ {
    //     for j := 0; j<ydim; j++ {
    //         grid[i][j][0] = minLongitude + float64(x_chunk)*float64(i)
    //         grid[i][j][1] = minLatidude + float64(y_chunk)*float64(j)
    //     }
    // }

    grid2D := make([][]int, xdim)
    for i := 0; i < xdim; i++ {
        grid2D[i] = make ([]int, ydim)
    }
    var queryWest float64
    var queryEast float64
    var querySouth float64
    var queryNorth float64

    switch ver {
    case "-v1":
        // YOUR SETUP CODE FOR PART 1


    case "-v2":
        // YOUR SETUP CODE FOR PART 2
    case "-v3":
        // YOUR SETUP CODE FOR PART 3
        // Part 3 - make a 2D array

        
        // Version 3: Step 1
        for k := 0; k<len(censusData); k++ {
            for i := 0; i<xdim-1;i++ {
                for j := 0; j<ydim-1; j++ {

                    queryWest = minLongitude + float64(x_chunk)*float64(i)
                    queryEast = minLongitude + float64(x_chunk)*float64(i+1)
                    querySouth = minLatidude + float64(y_chunk)*float64(j)
                    queryNorth = minLatidude + float64(y_chunk)*float64(j+1)

                    if censusData[k].longitude >= queryWest && censusData[k].longitude <= queryEast && censusData[k].latitude >= querySouth && censusData[k].latitude <= queryNorth {
                        grid2D[i][j] = grid2D[i][j] + censusData[k].population
                    }
                }
            }            
        }
        
        // Version 3: Step 1
        // for k := 0; k<len(censusData); k++ {
        //     for i := 0; i<xdim-1;i++ {
        //         for j := 0; j<ydim-1; j++ {
        //             if censusData[k].longitude >= grid[i][j][0] && censusData[k].longitude <= grid[i+1][j+1][0] && censusData[k].latitude >= grid[i][j][1] && censusData[k].latitude <= grid[i+1][j+1][1] {
        //                 grid2D[i][j] = grid2D[i][j] + censusData[k].population
        //                 // XTREEEEMME DEBUGGING SKEEELZ
        //                 // if censusData[k].population == 4213 {
        //                 //     fmt.Printf("%v%v%v%v\n", "longitude: ", censusData[k].longitude, " latitude: ", censusData[k].latitude) 
        //                 // }
        //             }
        //         }
        //     }            
        // }

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
        fmt.Printf("%v%v\n", "diff :: ",populationTotal - grid2D[xdim-1][0])

        // var percentage = (float64(grid2D[xdim-1][0])/float64(populationTotal)) * float64(100)
        // fmt.Printf("%v%v\n", "perage :: ",percentage)

        // XTREEEEMME DEBUGGING SKEEELZ SQUARED
        /*
        var random int = 0
        for i := 0; i<xdim;i++ {
            for j := 0; j<ydim; j++ {
                    fmt.Printf("%v%v%v%v%v\n", i, " ", j, " ", grid2D[i][j]) 
                    random = random + grid2D[i][j]
            }
        }    
        fmt.Printf("%v%v", "tot ",populationTotal)        
        fmt.Printf("%v%v", "ran ",random)        
        */


    case "-v4":
        // YOUR SETUP CODE FOR PART 4   
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

        west = west - 1     // ahh never mind
        south = south - 1   // ahh never mind
        // east = east - 1     // ahh never mind
        // north = north - 1   // ahh never mind

        var population int
        var percentage float64

        switch ver {
        case "-v1":
            // YOUR QUERY CODE FOR PART 1
            // for loop that runs over entire CensusData array, the IF condition inside checks which of the datapoints occur inside the "queried rectangle"...and adds their population
            // for i := 0; i<len(censusData); i++ {
            //     if censusData[i].longitude >= grid[west][south][0] && censusData[i].longitude < grid[east][north][0] && censusData[i].latitude >= grid[west][south][1] && censusData[i].latitude < grid[east][north][1]   {
            //         population += censusData[i].population
            //     }
            // }
            // percentage = (float64(population)/float64(populationTotal)) * float64(100)

        case "-v2":
            // YOUR QUERY CODE FOR PART 2
        case "-v3":
            // YOUR QUERY CODE FOR PART 3

            // var B,C,D int
            // if north-1 < 0 { B = 0 } else { B = grid2D[east][north-1] }  
            // if west-1 < 0 { C = 0 } else { C = grid2D[west-1][south] }
            // if west-1 < 0 || north-1 < 0 { D = 0 } else { D = grid2D[west-1][north-1] }
            // population = grid2D[east][south] - B - C + D
            // fmt.Printf("%v%v\n", "diff :: ", populationTotal - population)
            // percentage = (float64(population)/float64(populationTotal)) * float64(100)

            // population = grid2D[east][south] - grid2D[east][north-1] - grid2D[west-1][south] + grid2D[west-1][north-1]
            // bottom-right => grid2D[east][south]
            // bottom-left => grid2D[west][south]
            // top-right => grid2D[east][north]
            // top-left => grid2D[west][north]

            // Earlier Version...works fine
            if north-1 < 0 && west-1 >= 0 && south-1 >= 0{
                population = grid2D[east][south] - grid2D[west-1][south] + grid2D[west-1][south-1]
            } else if west-1 < 0 && north-1 >= 0{
                population = grid2D[east][south] - grid2D[east][north-1]
            } else if north-1 <0 && west-1 <0 && south-1 <0 {
                population = grid2D[east][south]
            } else {
                population = grid2D[east][south] - grid2D[east][north-1] - grid2D[west-1][south] + grid2D[west-1][south-1]  
            } 
            percentage = (float64(population)/float64(populationTotal)) * float64(100)


        case "-v4":
            // YOUR QUERY CODE FOR PART 4
        case "-v5":
            // YOUR QUERY CODE FOR PART 5
        case "-v6":
            // YOUR QUERY CODE FOR PART 6
        }

        fmt.Printf("%v %.2f%%\n", population, percentage)
    }
}

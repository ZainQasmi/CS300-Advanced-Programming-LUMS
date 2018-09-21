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


    grid2D := make([][]int, xdim)
    for i := 0; i < xdim; i++ {
        grid2D[i] = make ([]int, ydim)
    }



    switch ver {
    case "-v1":
        // YOUR SETUP CODE FOR PART 1

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

    case "-v2":
        // YOUR SETUP CODE FOR PART 2

    case "-v3":
        // YOUR SETUP CODE FOR PART 3
        // Part 3 - make a 2D array
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

        var x_chunk float64 = (maxLongitude - minLongitude) / float64(xdim)
        var y_chunk float64 = (maxLatitude - minLatidude) / float64(ydim) 

        var queryWest float64
        var queryEast float64
        var querySouth float64
        var queryNorth float64  

        // Version 3: Step 1
        for k := 0; k<len(censusData); k++ {
            for i := 0; i<xdim;i++ {
                for j := 0; j<ydim; j++ {

                    queryWest = minLongitude + float64(x_chunk)*float64(i)
                    queryEast = minLongitude + float64(x_chunk)*float64(i)+x_chunk
                    querySouth = minLatidude + float64(y_chunk)*float64(j)
                    queryNorth = minLatidude + float64(y_chunk)*float64(j)+y_chunk

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

        var population int
        var percentage float64

        switch ver {
        case "-v1":
            // YOUR QUERY CODE FOR PART 1
            west = west - 1     // ahh never mind
            south = south - 1   // ahh never mind

            var x_chunk float64 = (maxLongitude - minLongitude) / float64(xdim)
            var y_chunk float64 = (maxLatitude - minLatidude) / float64(ydim)

            var queryWest float64 = minLongitude + float64(x_chunk)*float64(west)
            var queryEast float64 = minLongitude + float64(x_chunk)*float64(east)
            var querySouth float64 = minLatidude + float64(y_chunk)*float64(south)
            var queryNorth float64 = minLatidude + float64(y_chunk)*float64(north)

            for i := 0; i<len(censusData); i++ {
                if censusData[i].longitude >= queryWest && censusData[i].longitude <= queryEast && censusData[i].latitude >= querySouth && censusData[i].latitude <= queryNorth {
                    population += censusData[i].population
                }
            }
            percentage = (float64(population)/float64(populationTotal)) * float64(100)

        case "-v2":
            // YOUR QUERY CODE FOR PART 2
        case "-v3":
            // YOUR QUERY CODE FOR PART 3

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

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

var populationTotal int
var maxLatitude float64
var minLatidude float64
var maxLongitude float64
var minLongitude float64

// var population int = 0
// var percentage float64
// var west, south, east, north int

func calRectangle (censusData []CensusGroup) () {

    populationTotal = 0
    maxLatitude = censusData[0].latitude
    minLatidude = censusData[0].latitude
    maxLongitude = censusData[0].longitude
    minLongitude = censusData[0].longitude

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

    return
}

func query (censusData [] CensusGroup, grid [][][] float64) {
	fmt.Printf("i iz here\n")
	for i := 0; i<len(censusData); i++ {
	    if censusData[i].longitude >= grid[west][south][0] && censusData[i].longitude < grid[east][north][0] && censusData[i].latitude >= grid[west][south][1] && censusData[i].latitude < grid[east][north][1]   {
	        population += censusData[i].population
	    }
	    // fmt.Printf("%v%v",  " ",i)
	}
	percentage = (float64(population)/float64(populationTotal)) * float64(100)
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
    // populationTotal := make(chan int)
    // maxLatitude := make(chan float64)
    // minLatidude := make(chan float64)
    // maxLongitude := make(chan float64)
    // minLongitude := make(chan float64)
    // var populationTotal, maxLatitude, minLatidude, maxLongitude, minLongitude = calRectangle(censusData);
    go calRectangle(censusData)
    // go func() {
    //     var _populationTotal int = 0
    //     var _maxLatitude float64 = censusData[0].latitude
    //     var _minLatidude float64 = censusData[0].latitude
    //     var _maxLongitude float64 = censusData[0].longitude
    //     var _minLongitude float64 = censusData[0].longitude

    //     for i := 0; i < len(censusData); i++ { // Can start with i=1 to reduce computation
    //             if censusData[i].latitude > _maxLatitude {
    //                 _maxLatitude = censusData[i].latitude
    //             } 
    //             if censusData[i].latitude < _minLatidude {
    //                 _minLatidude = censusData[i].latitude
    //             }
    //             if censusData[i].longitude > _maxLongitude {
    //                 _maxLongitude = censusData[i].longitude
    //             } 
    //             if censusData[i].longitude < _minLongitude {
    //                 _minLongitude = censusData[i].longitude
    //             }
    //             _populationTotal += censusData[i].population
    //     }

    //     // return populationTotal, maxLatitude, minLatidude, maxLongitude, minLongitude
    //     populationTotal<-_populationTotal
    //     maxLatitude<-_maxLatitude
    //     minLatidude<-_minLatidude
    //     maxLongitude<-_maxLongitude
    //     minLongitude<-_minLongitude
    // }()

    // <-populationTotal
    // <-maxLatitude
    // <-minLatidude
    // <-maxLongitude
    // <-minLongitude

    //     if censusData[i].latitude > maxLatitude {
    //         maxLatitude = censusData[i].latitude
    //     } 
    //     if censusData[i].latitude < minLatidude {
    //         minLatidude = censusData[i].latitude
    //     }
    //     if censusData[i].longitude > maxLongitude {
    //         maxLongitude = censusData[i].longitude
    //     } 
    //     if censusData[i].longitude < minLongitude {
    //         minLongitude = censusData[i].longitude
    //     }
    //     populationTotal += censusData[i].population
    // }

    fmt.Printf("%v%v\n", "xdim :: ",xdim);
    fmt.Printf("%v%v\n", "ydim :: ",ydim);
    fmt.Printf("%v%v\n", "populationTotal :: ", populationTotal);

    fmt.Printf("%v%v\n%v%v\n%v%v\n%v%v\n", "maxLatitude :: ",maxLatitude, "minLatitude :: ", minLatidude,"maxLongitude :: ",maxLongitude,"minLongitude :: ", minLongitude );

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
    case "-v6":
        // YOUR SETUP CODE FOR PART 6
    default:
        fmt.Println("Invalid version argument")
        return
    }

    for {
        // var west, south, east, north int
        n, err := fmt.Scanln(&west, &south, &east, &north)
        if n != 4 || err != nil || west<1 || west>xdim || south<1 || south>ydim || east<west || east>xdim || north<south || north>ydim {
            break
        }

        west = west - 1     // ahh never mind
        south = south - 1   // ahh never mind
        // east = east - 1
        // north = north - 1

        var population int = 0 //init with 0
        var percentage float64

        var x_chunk float64 = ((-minLongitude) - (-maxLongitude)) / float64(xdim-1)
        var y_chunk float64 = (maxLatitude - minLatidude) / float64(ydim-1)

        // fmt.Printf("%v%v\n", "xchunk :: ",minLongitude + float64(x_chunk)*float64(xdim));
        // fmt.Printf("%v%v\n", "ychunk :: ",minLatidude + float64(y_chunk)*float64(ydim));

        // var grid [100][50][2]float64

        // Initialize 3D Array of size xdim x ydim x 2
        // Last dimension has longitude in [0] and latitude in [1]
        grid := make([][][]float64, xdim)
        for i := 0; i < xdim; i++ {
            grid[i] = make ([][]float64, ydim)
            for j:= 0; j < ydim; j++ {
                grid[i][j] = make([]float64, 2)
            }
        }

        // Divide large rectangle into equal chunks of latitudes and longitudes
        // then store them in grid
        for i := 0; i<xdim;i++ {
            for j := 0; j<ydim; j++ {
                grid[i][j][0] = minLongitude + float64(x_chunk)*float64(i)
                grid[i][j][1] = minLatidude + float64(y_chunk)*float64(j)
            }
        }

        switch ver {
        case "-v1":
            // YOUR QUERY CODE FOR PART 1
        	go query(censusData, grid)

        	// fmt.Printf("i iz here\n")
        	// for i := 0; i<len(censusData); i++ {
        	//     if censusData[i].longitude >= grid[west][south][0] && censusData[i].longitude < grid[east][north][0] && censusData[i].latitude >= grid[west][south][1] && censusData[i].latitude < grid[east][north][1]   {
        	//         population += censusData[i].population
        	//     }
        	// }
        	percentage = (float64(population)/float64(populationTotal)) * float64(100)


            

        case "-v2":
            // YOUR QUERY CODE FOR PART 2
        case "-v3":
            // YOUR QUERY CODE FOR PART 3
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

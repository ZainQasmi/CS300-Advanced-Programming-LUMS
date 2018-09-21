var http = require("http");
var https = require("https");
const fs = require('fs')
const path = require('path')

const parse5 = require('parse5');
const file = fs.createWriteStream('exported.txt');
const parser = new parse5.SAXParser();

const writeFile = (file,data) => new Promise((resolve, reject) =>
    fs.writeFile(file, data, err => err ? reject(err) : resolve()))

const readFile = file => new Promise((resolve,reject) =>
	fs.readFile(file, (err,data) => err ? reject(err) : resolve(data)))

/*
var options = {
  hostname: 'stackoverflow.com',
  // hostname: 'lums.edu.pk',
  path: '/'
};
*/

var myList = [];
var myListDistinct = [];

const fetchHtml = oneURL => {
	return new Promise((resolve, reject) => {

		// if(oneURL.startsWith('https')) {
		//     let request = https.get(oneURL, function (res, err) {
		//     	console.log('err')
		//         res.pipe(parser)

		// 		parser.on('startTag', (arg1,arg2) => {
		// 			console.log(arg2)
		// 			if (arg1 == 'a') {
		// 				arg2.forEach ( tuple => {
		// 					console.log('here')
		// 					console.log(tuple.value)
		// 					if (tuple.name == 'href'){//} && tuple.value != '#') {
		// 							// console.log(tuple.value)
		// 							if (tuple.value[0] == '/' && tuple.value[1] != '/'){
		// 								// console.log(oneURL + tuple.value)
		// 								myList.push(oneURL + tuple.value)
		// 							}
		// 							else if (tuple.value[0] == '/' && tuple.value[1] == '/'){
		// 								// console.log(tuple.value)
		// 								myListDistinct.push('http:' + tuple.value)
		// 								myList.push('http:' + tuple.value)	
		// 							}
		// 							else if (tuple.value.substring(0,7) == 'http://' || tuple.value.substring(0,8) == 'https://'){
		// 								// console.log(tuple.value)
		// 								myList.push(tuple.value)
		// 							}
		// 					}		
		// 				})
		// 			}
		// 		});

		//         var data = "";
		//         res.on("data", buffer => data += buffer);
		//         res.on("end", function () {
		//             console.log("fetch complete hogaya");
		//             resolve(data);
		//         });
		//     });

		//     request.on('error', error => {
		//     	reject(new Error(error));
		//     });
		// } else {
		    let request = http.get(oneURL, function (res, err) {
		    	console.log(err)
		        res.pipe(parser)
		    	if (res.statusCode >= 100 && res.statusCode < 199) {
			        console.log(res.statusCode)
			    } else if (res.statusCode >= 300 && res.statusCode > 399) {
			        reject(new Error('Redirection'));
			    } else if (res.statusCode >= 400 && res.statusCode > 499) {
			        reject(new Error('Client errors'));
			    } else if (res.statusCode >= 500 && res.statusCode > 599) {
			        reject(new Error('Server error'));
			    }

				parser.on('startTag', (arg1,arg2) => {
					if (arg1 == 'a') {
						arg2.forEach ( tuple => {
							if (tuple.name == 'href'){//} && tuple.value != '#') {
									// console.log(tuple.value)
									if (tuple.value[0] == '/' && tuple.value[1] != '/'){
										// console.log(oneURL + tuple.value)
										myList.push(oneURL + tuple.value)
									}
									else if (tuple.value[0] == '/' && tuple.value[1] == '/'){
										// console.log(tuple.value)
										myListDistinct.push('http:' + tuple.value)
										myList.push('http:' + tuple.value)	
									}
									else if (tuple.value.substring(0,7) == 'http://' || tuple.value.substring(0,8) == 'https://'){
										// console.log(tuple.value)
										myList.push(tuple.value)
									}
							}		
						})
					}
				});

		        var data = "";
		        res.on("data", buffer => data += buffer);
		        res.on("end", function () {
		            console.log("fetch complete hogaya");
		            resolve(data);
		        });
		    });

		    request.on('error', error => {
		    	reject(new Error(error));
		    });
		// }


	});
};

siteInfo = {
    'rootDomain.com': {
        requestCount: 2,
        promisedDelay: 1000
    },
    'anotherDomain.com': {
        requestCount: 1,
        promisedDelay: 1000
    }
}

function fetch_parse_recursive (listOfUrls) {
	fetchHtml(listOfUrls)
	.then(retList => {
		retList.forEach ((myURL) =>{
			fetch_parse_recursive(myURL)
			// console.log(myURL)
		})
	}).catch(error => {
		console.error(error);
	})
}


readFile('config.json')
	.then(JSONdata => {
		let parsed = JSON.parse(JSONdata)
		// console.log(parsed.initialUrls)
		return parsed.initialUrls
	}).catch((error) => console.log(error))
	.then((JSONurls) => {
		JSONurls.forEach((oneURL) => {
				// console.log(oneURL)
				fetchHtml(oneURL)
				.then(data => {
					myList.forEach( (myURL)=> {
						console.log(myURL)
					})
					console.log("fetch pora complete hogaya")
				}).catch(error => {
					console.error(error);
				});
		})
	}).catch((error) => console.log(error))
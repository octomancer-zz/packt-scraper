package freelearning

type FreeLearningBooks struct {
	Count int                `json:"count"`
	Data  []FreeLearningBook `json:"data"`
}

type FreeLearningBook struct {
	ProductId string `json:"productId"`
	Summary   *FreeLearningBookSummary
}

type FreeLearningBookSummary struct {
	Title      string `json:"title"`
	Features   string `json:"features"`
	CoverImage string `json:"coverImage"`
	OneLiner   string `json:"oneLiner"`
}

/*

Example FreeLearningBooks

{"data":[{"id":"3c367558-6ba0-4464-b0c0-751fe15efb3f","productId":"9781788629515","availableFrom":"2019-02-19T00:00:00.000Z","expiresAt":"2019-02-20T00:00:00.000Z","limitedAmount":false,"amountAvailable":null,"details":null,"priority":0,"createdAt":"2019-02-15T11:20:03.794Z","updatedAt":"2019-02-15T11:20:54.733Z","deletedAt":null}],"count":1}

Example FreeLearningBookSummary

{"title":"Learning SciPy for Numerical and Scientific Computing","type":"books","coverImage":"https://d255esdrn735hr.cloudfront.net/sites/default/files/1622OS.jpg","productId":"9781782161622","isbn13":"9781782161622","oneLiner":"For solving complex problems in mathematics, science, or engineering, SciPy is the solution. Building on your basic knowledge of Python, and using a wealth of examples from many scientific fields, this book is your expert tutor.","pages":150,"publicationDate":"2013-02-22T05:10:00.000Z","length":"4 hours 30 minutes","about":"<p>It's essential to incorporate workflow data and code from various sources in order to create fast and effective algorithms to solve complex problems in science and engineering. Data is coming at us faster, dirtier, and at an ever increasing rate. There is no need to employ difficult-to-maintain code, or expensive mathematical engines to solve your numerical computations anymore. SciPy guarantees fast, accurate, and easy-to-code solutions to your numerical and scientific computing applications.<br /><br />\"Learning SciPy for Numerical and Scientific Computing\" unveils secrets to some of the most critical mathematical and scientific computing problems and will play an instrumental role in supporting your research. The book will teach you how to quickly and efficiently use different modules and routines from the SciPy library to cover the vast scope of numerical mathematics with its simplistic practical approach that's easy to follow.<br /><br />The book starts with a brief description of the SciPy libraries, showing practical demonstrations for acquiring and installing them on your system. This is followed by the second chapter which is a fun and fast-paced primer to array creation, manipulation, and problem-solving based on these techniques.<br /><br />The rest of the chapters describe the use of all different modules and routines from the SciPy libraries, through the scope of different branches of numerical mathematics. Each big field is represented: numerical analysis, linear algebra, statistics, signal processing, and computational geometry. And for each of these fields all possibilities are illustrated with clear syntax, and plenty of examples. The book then presents combinations of all these techniques to the solution of research problems in real-life scenarios for different sciences or engineering — from image compression, biological classification of species, control theory, design of wings, to structural analysis of oxides.</p>","learn":"<ul>\r\n<li>Learn to store and manipulate large arrays of data in any dimension</li>\r\n<li>Accurately evaluate any mathematical function in any given dimension, as well as its integration, and solve systems of ordinary differential equations with ease</li>\r\n<li>Learn to deal with sparse data to perform any known interpolation, extrapolation, or regression scheme on it</li>\r\n<li>Perform statistical analysis, hypothesis test design and resolution, or data mining at high level, including clustering (hierarchical or through vector quantization), and learn to apply them to real-life problems</li>\r\n<li>Get to grips with signal processing — filtering audio, images, or video to extract information, features, or removing components</li>\r\n<li>Effectively learn about window functions, filters, spectral theory, LTY systems theory, morphological operations, and image interpolation</li>\r\n<li>Acquaint yourself with the power of distances, Delaunay triangulations, and Voronoi diagrams for computational geometry, and apply them to various engineering problems</li>\r\n<li>Wrap code in other languages directly into your SciPy-based workflow, as well as incorporating data written in proprietary format (audio or image, for example), or from other software suites like Matlab/Octave</li>\r\n</ul>","features":"<ul>\r\n<li>Perform complex operations with large matrices, including eigenvalue problems, matrix decompositions, or solution to large systems of equations</li>\r\n<li>Step-by-step examples to easily implement statistical analysis and data mining that rivals in performance any of the costly specialized software suites</li>\r\n<li>Plenty of examples of state-of-the-art research problems from all disciplines of science, that prove how simple, yet effective, is to provide solutions based on SciPy</li>\r\n</ul>\r\n<p>&nbsp;</p>","authors":["11440"],"shopUrl":"/big-data-and-business-intelligence/learning-scipy-numerical-and-scientific-computing","readUrl":"/book/big-data-and-business-intelligence/9781782161622","category":"big-data-and-business-intelligence","earlyAccess":false,"available":true}

*/

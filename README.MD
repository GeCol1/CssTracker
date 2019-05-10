#  User Tracking POC 

This code is from this [Gist](https://gist.github.com/wybiral/c8f46fdf1fc558d631b55de3a0267771).

Run the code and open 2 pages: on "/" and on "/watch".

# Explanation

The generated page has a div with a :hover tag on each cell. This tag call an image and the server intercept this. As there is a grid, the server knows where the mouse was.
It should be interested to see how the browser cache handles this.
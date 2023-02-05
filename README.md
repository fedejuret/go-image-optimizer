# Image Optimizer

This project is the result of an internship I'm doing to learn how to use Go.
Fortunately, I have been able to put it to real use in the company where I work, and many colleagues have benefited from this tool.

## How does it work?
In order to develop the CLI, I have used the [Cobra](https://github.com/spf13/cobra) library, which allows you to make console applications very easily.
Then, to compress the images I have used [Bimg](https://github.com/h2non/bimg)

## Start using it
Simply clone the repository and build using `go build` inside the project. This command will generate a file called `image-optimizer` which is an executable that is ready to start using.
**Personally I have tested it ONLY on Linux Ubuntu, so I can't confirm if it works on Windows**

## To collaborate
To be able to collaborate, simply fork the project and then send your pull requests, I am open to suggestions, improvements, and especially advice on Go, since it is a new language for me.
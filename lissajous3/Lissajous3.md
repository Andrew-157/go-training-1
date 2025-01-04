# Lissajous3

Check out its usage, pretty cool:

```commandline
root@fedora:~/go_dir/lissajous3# go run .
Generate a lissajous gif with custom colors

Choose background color of the image (enter a number or color name, enter 'exit' to exit the program)
black(0), white(1), red(2), green(3), blue(4): 1
You chose white(1) background

Choose primary color of the image (enter a number or color name, enter 'exit' to exit the program)
black(0), white(1), red(2), green(3), blue(4): 1 
You chose white(1) primary

Background and Primary colors are the same, are you sure?[yes y no n] n
Choose primary color of the image (enter a number or color name, enter 'exit' to exit the program)
black(0), white(1), red(2), green(3), blue(4): exit
Exiting...
root@fedora:~/go_dir/lissajous3# go run .
Generate a lissajous gif with custom colors

Choose background color of the image (enter a number or color name, enter 'exit' to exit the program)
black(0), white(1), red(2), green(3), blue(4): 4
You chose blue(4) background

Choose primary color of the image (enter a number or color name, enter 'exit' to exit the program)
black(0), white(1), red(2), green(3), blue(4): 0
You chose black(0) primary

You chose blue and black for the bg and pr, respectively

Enter a filename(relative or full path) to where gif will be written to(enter 'exit' to exit the program): ./
Invalid filename './', filename must end with .gif: /
Invalid filename '/', filename must end with .gif: mygif.gif
Creating a new file: mygif.gif

Your gif was generated!
root@fedora:~/go_dir/lissajous3# go run .
Generate a lissajous gif with custom colors

Choose background color of the image (enter a number or color name, enter 'exit' to exit the program)
black(0), white(1), red(2), green(3), blue(4): 3
You chose green(3) background

Choose primary color of the image (enter a number or color name, enter 'exit' to exit the program)
black(0), white(1), red(2), green(3), blue(4): 0
You chose black(0) primary

You chose green and black for the bg and pr, respectively

Enter a filename(relative or full path) to where gif will be written to(enter 'exit' to exit the program): mygif.gif
Using existing file: mygif.gif

Your gif was generated!
root@fedora:~/go_dir/lissajous3# go run .
Generate a lissajous gif with custom colors

Choose background color of the image (enter a number or color name, enter 'exit' to exit the program)
black(0), white(1), red(2), green(3), blue(4): 0
You chose black(0) background

Choose primary color of the image (enter a number or color name, enter 'exit' to exit the program)
black(0), white(1), red(2), green(3), blue(4): 1
You chose white(1) primary

You chose black and white for the bg and pr, respectively

Enter a filename(relative or full path) to where gif will be written to(enter 'exit' to exit the program): exit
Exiting...
```

# ascii-art

## Objectives

Ascii-art consists on receiving a `string` as an argument and outputting the `string` in a graphic representation of ASCII.

- This project should handle numbers, letters, spaces, special characters and `\n`.
- Take a look at the ASCII manual.

This project will help you learn about:

- Client utilities.
- The Go file system(**fs**) API.
- Ways to receive data.
- Ways to output data.
- Manipulation of strings.
- Manipulation of structures.

## Instructions

- Your project must be written in **Go**.
- The code must respect the [**good practices**](https://public.01-edu.org/subjects/good-practices.en).
- It is recommended that the code present a **test file**.

- Some **banner** files with a specific graphical template representation using ASCII will be given. The files are formatted in a way that is not necessary to change them.

  - [shadow](shadow.txt)
  - [standard](standard.txt)
  - [thinkertoy](thinkertoy.txt)

## Banner Format

- Each character has an height of 8 lines.
- Characters are separated by a new line `\n`.
- Here is an example of ' ', '!' and '"'(one dot represents one space) :

```console

......
......
......
......
......
......
......
......

._..
|.|.
|.|.
|.|.
|_|.
(_).
....
....

._._..
(.|.).
.V.V..
......
......
......
......
......

etc
```

## Usage

```console
$ go run . "hello" | cat -e
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $                           
```

```console
$ go run . "HeLlO" | cat -e
 _    _          _        _    ____   $
| |  | |        | |      | |  / __ \  $
| |__| |   ___  | |      | | | |  | | $
|  __  |  / _ \ | |      | | | |  | | $
| |  | | |  __/ | |____  | | | |__| | $
|_|  |_|  \___| |______| |_|  \____/  $
                                      $
                                      $
```

```console
$ go run . "Hello There" | cat -e
 _    _          _   _                 _______   _                           $
| |  | |        | | | |               |__   __| | |                          $
| |__| |   ___  | | | |   ___            | |    | |__     ___   _ __    ___  $
|  __  |  / _ \ | | | |  / _ \           | |    |  _ \   / _ \ | '__|  / _ \ $
| |  | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/ $
|_|  |_|  \___| |_| |_|  \___/           |_|    |_| |_|  \___| |_|     \___| $
                                                                             $
                                                                             $
```

```console
$ go run . "1Hello 2There" | cat -e
     _    _          _   _                         _______   _                           $
 _  | |  | |        | | | |                ____   |__   __| | |                          $
/ | | |__| |   ___  | | | |   ___         |___ \     | |    | |__     ___   _ __    ___  $
| | |  __  |  / _ \ | | | |  / _ \          __) |    | |    |  _ \   / _ \ | '__|  / _ \ $
| | | |  | | |  __/ | | | | | (_) |        / __/     | |    | | | | |  __/ | |    |  __/ $
|_| |_|  |_|  \___| |_| |_|  \___/        |_____|    |_|    |_| |_|  \___| |_|     \___| $
                                                                                         $
                                                                                         $
```

```console
$ go run . "{Hello There}" | cat -e
   __  _    _          _   _                 _______   _                           __    $
  / / | |  | |        | | | |               |__   __| | |                          \ \   $
 | |  | |__| |   ___  | | | |   ___            | |    | |__     ___   _ __    ___   | |  $
/ /   |  __  |  / _ \ | | | |  / _ \           | |    |  _ \   / _ \ | '__|  / _ \   \ \ $
\ \   | |  | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/   / / $
 | |  |_|  |_|  \___| |_| |_|  \___/           |_|    |_| |_|  \___| |_|     \___|  | |  $
  \_\                                                                              /_/   $
                                                                                         $
```

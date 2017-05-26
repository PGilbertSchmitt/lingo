# Lingo - Phonology Generator tool

## Concept

Lingo is intended as a conlanger's tool. Given a phonological inventory and set of rules, Lingo can produce a random list of words. While morphology is still a while away for this project, any conlanger can come up with a reasonable lexicon with minimal effort needed for tweaking, even if it's just for inspiration.

## Method

Lingo will accept a text file of configuration, with specific formatting (not yet implemented). This configuration will include the phonetic inventory (all phonemes (distinct sounds) that make up that language), syllable composition, and replacement rules.

Using the inventory and composition, lingo will generate as many

## Golang

Lingo is written in Golang, for a few reasons:

1. It's lightning fast

Lingo is intended to be run as a command line tool, though eventually it will be accessible through a web application, and so it needs to have a modicum of haste. There's likely a lot of processing that needs to be done, and the built-in multithreading will be very friendly.

2. Readability / Ease of use

While C/C++ would have worked (eventually), there are many aspects of string manipulation, array handing, and archaic syntax that I refuse to deal with.

3. I like it

Will you get off my back, Alec?

## Input File

Lingo accepts a single argument at the moment, the file containing the rules for how syllables are built. The format is specific, but flexible. An example can be found in [this example file](example-input.txt).

The file should have three headers, all in uppercase and preceded by an octathorpe (#) and a space ( ). Every line after the header until a space or EOF is considered to be part of that section. The three headers are:

- Syllables
- Consonsants
- Vowels

#### SYLLABLES

All lines after `# SYLLABLES` should be of the common form:

```Text
cccvccc
vvcc
cv
```

It accepts any valid number of c's, v's, and more c's. As the interpreter is very simple, if there are forms such as `ccvvccvv`, it will not work as intended. There should only be one set of v's.

#### CONSONANTS

A list of consonants and their weights. Each consonant should be in it's IPA character representation like so

```Text
t 10
p 20
ʁ 15
ɸ 5
```

#### VOWELS

Exactly the same as consonants.

```Text
a 16
e 5
```

### Replacement Rules

Eventually, the file will also include replacement rules for generated word strings. I'm sure these rules exist in linguistics proper, but I'm specifically using a rule format similar to the Sound Change Applier found here: http://www.zompist.com/sca2.html

The way we're designing this is probably going to be a little more complex, but should give you better control over your output.

## Features to implement

Very broadly: the roadmap looks thusly:

[x] Accept configuration file  
[ ] Generate initial sequences  
[ ] Apply rules  
[ ] Generate words  

## hardcode.rb

There are two text files and a ruby script in the `misc` directory. The text files are all recognized pulmonic consonants and all vowels, according the the IPA chart on Wikipedia. The Ruby script takes those two files and generates the `phonebank.go` file. If you want to tweak either of the text files to add your own consonants, feel free and run the script. Just understand that the current implementation does not allow for any non-essential phone attributes such as stress, intonation, and aspiration. I wanted to develop a minimal working version before getting into the parts of phonology that scare me the most.

## Issues

I am far from an expert in linguistics. I know the very basics of phonology, and next to nothing else of the topic. I am building this project with the help of several linguists, a dash of Wikipedia, and a pinch of intuition. It will not be perfect, and many refinements have been made before the first line of code was even written (there will certainly be many more to come). However, I'm confident that I can get a working version out soon. Please, if any of this seems counterintuitive or factually wrong, I would love to know about it. Leave an issue, or send me an email at PGilbertSchmitt@gmail.com
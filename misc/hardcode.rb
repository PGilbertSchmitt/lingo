# This script builds a hard coded list of all phones
# taken from the consonants.txt and vowels.txt files.
# This approach is taken to trade time for space and
# make this program as fast as possible. I don't feel
# bad because I didn't have to write it by hand, just
# the script that does.

require 'csv'

go_list = File.open("../phonology/phonebank.go", "w")

go_list.write("// File generated by 'misc/hardcode.rb'\n\n")
go_list.write("package phonology\n\n")

# Consonants
go_list.write("// AllConsonants returns all reasonable pulmonic consonants\n")
go_list.write("func AllConsonants() []Consonant {\n")
go_list.write("\treturn []Consonant{\n")

CSV.foreach("./consonants.txt") do |row|
    point = row[0].strip
    method = row[1].strip
    voiced = row[2].strip == 'T' ? 1 : 2
    code = row[3].strip

    go_list.write("\t\tNewConsonant(\n")
    go_list.write("\t\t\t\'#{code}\',\n")
    go_list.write("\t\t\t#{voiced},\n")
    go_list.write("\t\t\t#{point},\n")
    go_list.write("\t\t\t#{method},\n")
    go_list.write("\t\t),\n")
end

# Wrapping up consonants
go_list.write("\t}\n}\n\n")

# Vowels
go_list.write("// AllVowels returns all reasonable vowels\n")
go_list.write("func AllVowels() []Vowel {\n")
go_list.write("\treturn []Vowel{\n")

CSV.foreach("./vowels.txt") do |row|
    openness = row[0].strip
    frontness = row[1].strip
    rounded = row[2].strip == 'T' ? 1 : 2
    code = row[3].strip

    go_list.write("\t\tNewVowel(\n")
    go_list.write("\t\t\t\'#{code}\',\n")
    go_list.write("\t\t\t#{rounded},\n")
    go_list.write("\t\t\t#{frontness},\n")
    go_list.write("\t\t\t#{openness},\n")
    go_list.write("\t\t),\n")
end

# Wrapping up vowels
go_list.write("\t}\n}")

go_list.close()
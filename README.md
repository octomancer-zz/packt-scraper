# packt-scraper


!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

This is a beta script that does what I need it to do. I'm providing it as a starting point for others. As long as "BOOK_DIR", "HTML_DIR" and "INFO_DIR" point to empty directories, it's not going to mess anything up. It won't try and download any book files unless you give the --fetch (-f) flag. Good luck! Run with "-v 1" to see increased logging.

!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

# INSTALL

Clone the repo, get deps with "go get -v" and install with "go install". Then you should have a "packt-scraper" binary in your path.

# CONFIGURATION FILES

You'll need a config file in your home directory called ".packt-scraper.yaml":

    # login details for packtpub.com
    email: YOUR_EMAIL_ADDRESS
    password: YOUR_PASSWORD
    # where to save the ebooks
    bookdir: BOOK_DIR
    # where to save any html files (can be same as above)
    htmldir: HTML_DIR
	# where to save book JSON
	infodir: INFO_DIR
    # who gets the emails
    recipients:
      - EMAIL1
      - EMAIL2

You'll also need a config file called ".mail.yaml" which configures an outgoing mail server to send the daily email. This works for gmail, YMMV:

    mailhost: smtp.gmail.com:587
    maildomain: smtp.gmail.com
    mailuser: YOUR_GMAIL_ADDRESS
    mailpass: YOUR_GMAIL_PASSWORD

# IMPORTANT

This script looks in "BOOK_DIR" as set in the config file to see if you already have a book downloaded. It does this by looking for a file that starts with the 11 digit productId and ends with the extension. Depending on how you name your book files, you may need to tweak the function "func FillBookFiles()" in locations.go for your use case.

# USAGE

You can run this to get basic help:

`$ packt-scraper help`

and you can run

`$ packt-scraper help COMMAND`

to get help for any of the subcommands

The first thing you'll need to do is download the JSON for the books you've claimed:

`$ packt-scraper fetchMyEbooks`

To get the daily free offer and send an email about it:

`$ packt-scraper freeLearning

To see a list of already-downloaded books the script has been able to recognise:

`$ packt-scraper show`

MailHog 
=========

MailHog is an email testing tool for developers:

* Configure your application to use MailHog for SMTP delivery
* View messages in the web UI, or retrieve them with the JSON API
* Optionally release messages to real SMTP servers for delivery

### Configuration

Check out how to [configure MailHog](/docs/CONFIG.md), or use the default settings:
  * the SMTP server starts on port 1025
  * the HTTP server starts on port 8025
  * in-memory message storage

### Features

See [MailHog libraries](docs/LIBRARIES.md) for a list of MailHog client libraries.

* ESMTP server implementing RFC5321
* Support for SMTP AUTH (RFC4954) and PIPELINING (RFC2920)
* Web interface to view messages (plain text, HTML or source)
  * Supports RFC2047 encoded headers
* Real-time updates using EventSource
* Release messages to real SMTP servers
* Chaos Monkey for failure testing
  * See [Introduction to Jim](/docs/JIM.md) for more information
* HTTP API to list, retrieve and delete messages
  * See [APIv1](/docs/APIv1.md) and [APIv2](/docs/APIv2.md) documentation for more information
* [HTTP basic authentication](docs/Auth.md) for MailHog UI and API
* Multipart MIME support
* Download individual MIME parts
* In-memory message storage
* MongoDB and file based storage for message persistence
* Lightweight and portable
* No installation required

### TODO

* Clean up old scripts and READMEs in the packages since they used to be independent repositories
* Improve documentation within the code and linting
* Ensure API docs are up-to-date
* Improve tests
* Reduce memory footprint
* Use `embed` for the UI instead of generating byte strings

#### sendmail

[mhsendmail](./sendmail) is a sendmail replacement for MailHog.

It redirects mail to MailHog using SMTP.

You can also use `MailHog sendmail ...` instead of the separate mhsendmail binary.

Alternatively, you can use your native `sendmail` command by providing `-S`, for example:

```bash
/usr/sbin/sendmail -S mail:1025
```

For example, in PHP you could add either of these lines to `php.ini`:

```
sendmail_path = /usr/local/bin/mhsendmail
sendmail_path = /usr/sbin/sendmail -S mail:1025
```

#### Web UI

![Screenshot of MailHog web interface](/docs/MailHog.png "MailHog web interface")

### Project History

This project was originally forked from the github.com/mailhog/MailHog repository after it appeared to be no longer maintained, 
while exhibiting critical flaws. I have decided to take over this project to ensure it receives regular maintenance for 
the benefit of the larger Go community. Since taking over, Mailhog has been transformed into a monorepo with simplified package
names and with support for Go Modules.

I'd like to thank Ian Kent for his hard work on the original iteration of the application.

### Licence

Currently Copyrighted ©‎ 2021, Scott Finlay

Initial Copyright ©‎ 2014 - 2017, Ian Kent (http://iankent.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.

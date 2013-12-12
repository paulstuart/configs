configs
=======

Encode and decode encrypted config files for server credentials.

This is intended to be used in conjunction with another app that will actually use the credentials.
It is separate so that the config files can be decoded, updated and encoded, but the the other app
can only decode internally.

Because it uses an embedded salt string and private key, the two need to be the same between both apps.

To change the default salt string without modifying the code:

go build/test/install -ldflags "-X secrets.salty my-new-salt-string"

## 0.2.1 (June 20, 2019)

* fix gpg_encrypted_message updates issue with Terraform 0.12.x
* use sensitive for content field to make sure it does not leak
* drop using SchemaStateFunc for result, as it has no effect with Terraform 0.12.x
* make gpg_encrypted_message resource use sha256 of content as ID

## 0.2.0 (June 15, 2019)

* Add Terraform 0.12 compatibility
* Restructure code to standard layout
* Add Makefile to document common tasks

## 0.1.0 (April 12, 2019)

* Initial release

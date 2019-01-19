# README

When Terraforming this stack, you need to target the SQL module first, then everything else.

``` shell
terraform apply --auto-approve --target=module.sql
...
terraform apply --auto-approve
```

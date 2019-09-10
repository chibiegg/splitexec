# split and exec()

split stdin into parts and execute command each part.

## Usage

### Example

```bash
some output | splitexec -s 100000000000 rclone rcat -v s3:bucket/key.%05d
```

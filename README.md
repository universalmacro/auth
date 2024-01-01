# auth-api

authorization 服務，登入，創建帳號等等服務  
| UAT |PRD |
|---|----|
| https://auth-uat.sum-foods.com | https://auth.sum-foods.com |

# APIs

詳細請看 openapi.yaml

```bash
openapi-generator generate -g typescript-fetch -i ./openapi.yaml -o typescript-fetch --additional-properties=npmName=@universalmacro/auth-ts-sdk
npm publish --access=public
```

```bash
openapi-generator generate -g dart -i ./openapi.yaml -o dart-client
```

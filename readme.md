Commands to run and deploy are:

- mkdir golang_vercel/
- cd golang_vercel/
- go mod init myapp
- go get github.com/gin-gonic/gin
- Fill code in folder accordingly

- git add . && git commit -m "C" && git push origin main && vercel . && vercel --prod
  (or)
  git add . && git commit -m "C" && git push origin main && vercel --prod

- URL to check : https://golang-vercel-ten.vercel.app/api

- Reference:
  - https://github.com/gin-gonic/gin
  - https://www.youtube.com/watch?v=jsdyAWc_x08
  - https://github.com/bobwatcherx/goDeployVercel

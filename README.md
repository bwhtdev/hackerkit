# hackerkit
> A SaaS Boilerplate for Extreme **Effiency** and **Eloquence**

### Stack
- Deployment
  - Docker/Docker Compose
- Reverse proxy
  - Nginx
- Database
  - Postgres
- Backend
  - Go
  - Air CLI
- Frontend
  - Astro
  - Svelte
  - Tailwind CSS
  - AlpineJS
  - TypeScript

### TODO list of proposed features
- Implement ChatGPT 3.5 API in backend
- Implement hotkeys/shortcuts on btns [also have them marked] (context menu too?)
- Add email API
- Deploy to Digital Ocean or another cheap VPS
- Learn K8s?
- Add thirdparty OAuth
- Add dark mode (easy)
- Implement Stripe payment API

### Command to prepare env vars
```
touch .env
ln .env frontend/.env
ln .env backend/.env
touch db/password.txt
```

### Commands to run (locally)
```
# make postgres db
docker run postgres...

# run backend
go run backend/main.go

# run frontend
cd frontend && pnpm i && pnpm start
```

### Commands to run (on deployment)
```
docker compose up --build -d
```

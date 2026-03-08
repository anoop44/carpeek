# Stage 1: Build
FROM node:20-alpine AS builder
WORKDIR /app

# Declare build arguments for NEXT_PUBLIC_* env vars
# These MUST be available at build time because Next.js inlines them into the client JS bundle
ARG NEXT_PUBLIC_API_URL
ARG NEXT_PUBLIC_GOOGLE_CLIENT_ID

# Make them available as env vars during the build
ENV NEXT_PUBLIC_API_URL=$NEXT_PUBLIC_API_URL
ENV NEXT_PUBLIC_GOOGLE_CLIENT_ID=$NEXT_PUBLIC_GOOGLE_CLIENT_ID

COPY package.json package-lock.json ./
RUN npm ci
COPY . .

RUN npm run build

# Stage 2: Run
FROM node:20-alpine AS runner
WORKDIR /app
ENV NODE_ENV=production
# Next.js collects completely anonymous telemetry data about general usage.
# Learn more here: https://nextjs.org/telemetry
# Uncomment the following line in case you want to disable telemetry during runtime.
ENV NEXT_TELEMETRY_DISABLED=1

COPY --from=builder /app/package.json /app/package-lock.json ./
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public
COPY --from=builder /app/node_modules ./node_modules

EXPOSE 3000

# Set hostname to 0.0.0.0 to be able to access from outside the container
ENV HOSTNAME="0.0.0.0"
ENV PORT=3000

CMD ["npx", "next", "start", "-H", "0.0.0.0"]

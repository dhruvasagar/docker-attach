FROM node:latest

COPY . .
RUN npm install

ENV NODE_INSPECT_RESUME_ON_START 1

EXPOSE 1234

CMD ["npm", "start"]

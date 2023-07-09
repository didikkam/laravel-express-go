const express = require("express");
const { PrismaClient } = require("@prisma/client");
const prisma = new PrismaClient();

const app = express();
app.get("/", (req, res) => {
  res.send("Halo, dunia!");
});
app.get("/users", async (req, res) => {
  const users = await prisma.users.findMany();
  res.json(users);
});

const port = 3000;
app.listen(port, () => {
  console.log(`Aplikasi berjalan pada http://localhost:${port}`);
});

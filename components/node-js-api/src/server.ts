import fastify from "fastify";
import { Logger } from "./logger";

const server = fastify();
const logger = Logger.getInstance();

server.get("/", async (request, reply) => {
  return { message: "Hello, World! from my Node App!!" };
});

const gracefulShutdown = async (signal: string) => {
  logger.info({
    action: "server",
    message: `Recebido sinal ${signal}. Encerrando o servidor...`,
    data: {},
  });

  try {
    await server.close();
    logger.info({
      action: "server",
      message: "Servidor encerrado com sucesso.",
      data: {},
    });
    process.exit(0);
  } catch (err) {
    logger.error({
      action: "server",
      message: "Erro ao encerrar o servidor.",
      data: { error: err },
    });
    process.exit(1);
  }
};

process.on("SIGINT", () => gracefulShutdown("SIGINT"));
process.on("SIGTERM", () => gracefulShutdown("SIGTERM"));

export { server };

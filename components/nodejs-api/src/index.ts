import "dotenv/config";
import { server } from "./server";
import { Logger } from "./logger";

const logger = Logger.getInstance();
const port = process.env.PORT ? Number(process.env.PORT) : 3000;

server.listen({ port }, (err, address) => {
  logger.info({
    action: "index",
    message: `Servidor rodando`,
    data: { address, port },
  });
  if (err) {
    logger.error({
      action: "index",
      message: "Erro ao iniciar o servidor.",
      data: { error: err },
    });
    process.exit(1);
  }
});

// 起動時にDBなどへの外部接続ができるかどうかの確認をする
import { PrismaClient } from "@prisma/client";
import logger from "../helpers/logger.js";

export async function PreCheck(prisma: PrismaClient) {
  try {
    logger.info("Pre boot check");
    await prisma.$connect();
    logger.info("Check succeed");
  } catch (e) {
    logger.fatal(e);
    process.exit(1);
  }
}

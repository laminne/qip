import { Prisma } from "@prisma/client";
import {
  AlreadyExistsError,
  InternalError,
  InvalidValueError,
  NotFoundError,
} from "../../helpers/errors.js";

// Prismaのエラーを汎用エラーに変換
export function PrismaErrorConverter(e: unknown): Error {
  if (e instanceof Prisma.PrismaClientKnownRequestError) {
    // cf. https://www.prisma.io/docs/reference/api-reference/error-reference
    switch (e.code) {
      case "P2001":
        return new NotFoundError();
      case "P2002":
        return new AlreadyExistsError();
      case "P2007":
        return new InvalidValueError();
      case "P2025":
        return new NotFoundError();
      default:
        return new InternalError();
    }
  }
  return new InternalError();
}

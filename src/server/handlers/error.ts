// InternalなエラーをAPIのエラーに変換
export function ErrorConverter(e: Error): [number, { message: string }] {
  switch (e.name) {
    case "InternalError":
      return [500, { message: "Internal error" }];
    case "NotFoundError":
      return [404, { message: "Not found" }];
    case "InvalidValueError":
      return [400, { message: "Request value is not valid" }];
    case "AlreadyExistsError":
      return [400, { message: "Already exists" }];
    default:
      return [500, { message: "Internal error" }];
  }
}

// オブジェクトが見つからなかった場合
export class NotFoundError extends Error {
  static {
    this.prototype.name = "NotFoundError";
  }
}

// 内部のエラー
export class InternalError extends Error {
  static {
    this.prototype.name = "InternalError";
  }
}

// バリデーション失敗/不正な値のとき
export class InvalidValueError extends Error {
  static {
    this.prototype.name = "InvalidValueError";
  }
}

// すでにオブジェクトが存在するとき
export class AlreadyExistsError extends Error {
  static {
    this.prototype.name = "AlreadyExistsError";
  }
}

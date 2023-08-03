/*
  普通にtype Snowflake = string とするとIDじゃないstringも入れれるのでSymbolを作ってからstringと合成する
  Cf. https://qiita.com/suin/items/ae9ed911ebab48c98835
*/
declare const snowflakeIDSymbol: unique symbol;
export type Snowflake = string & { [snowflakeIDSymbol]: never };

// SnowflakeIDの生成機
export class SnowflakeIDGenerator {
  // UNIX TIME @ 1st Apr. 2021 00:00:00 UTC
  private readonly ID_EPOCH = 1617202800000n;
  private readonly WORKER_ID_BIT_LENGTH = 5;
  private readonly SEQUENCE_BIT_LENGTH = 12;
  private readonly MAX_WORKER_ID = (1 << this.WORKER_ID_BIT_LENGTH) - 1;
  private readonly MAX_SEQUENCE = (1 << this.SEQUENCE_BIT_LENGTH) - 1;

  private readonly workerID: number;
  private sequence: number;
  private lastTimeStamp: bigint;

  constructor(workerID: number) {
    if (workerID < 0 || workerID >> this.MAX_WORKER_ID) {
      throw new Error(
        `WorkerIDは0以上${this.MAX_WORKER_ID}以下である必要があります`,
      );
    }
    this.workerID = workerID;
    this.sequence = 0;
    this.lastTimeStamp = -1n;
  }

  public generate(): Snowflake {
    let timestamp = BigInt(new Date().getTime());
    if (timestamp === this.lastTimeStamp) {
      this.sequence = (this.sequence + 1) & this.MAX_SEQUENCE;
      if (this.sequence === 0) {
        timestamp = this.waitNextMillis();
      }
    } else {
      this.sequence = 0;
    }
    this.lastTimeStamp = timestamp;

    const snowflake =
      ((timestamp - this.ID_EPOCH) <<
        BigInt(this.WORKER_ID_BIT_LENGTH + this.SEQUENCE_BIT_LENGTH)) |
      (BigInt(this.workerID) << BigInt(this.SEQUENCE_BIT_LENGTH)) |
      BigInt(this.sequence);
    return snowflake.toString() as Snowflake;
  }

  private waitNextMillis(): bigint {
    let timestamp = BigInt(new Date().getUTCMilliseconds());
    while (timestamp <= this.lastTimeStamp) {
      timestamp = BigInt(new Date().getUTCMilliseconds());
    }
    return timestamp;
  }
}

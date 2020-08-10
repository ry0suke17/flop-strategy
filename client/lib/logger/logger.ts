export class Logger {
  static init() {}

  static setUserId(id: string) {}

  static async error(err: any) {
    console.error(err);
    await this.captureException(err);
  }

  static async captureException(err: any) {
    //TODO: Sentryなどにログを送るなどしたらよりいい。
  }
}

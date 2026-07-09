enum LogLevel {
  INFO = "info",
  WARN = "warn",
  ERROR = "error",
  DEBUG = "debug",
}

type LogFormat = {
  timestamp: string;
  level: LogLevel;
  action: string;
  message: string;
  data: Record<string, unknown>;
};

export class Logger {
  private static instance: Logger;

  private constructor() {}

  public static getInstance(): Logger {
    if (!Logger.instance) {
      Logger.instance = new Logger();
    }
    return Logger.instance;
  }

  private log(data: Omit<LogFormat, "timestamp">): void {
    console.log(
      JSON.stringify({
        timestamp: new Date().toISOString(),
        level: data.level,
        action: data.action,
        message: data.message,
        data: data.data,
      }),
    );
  }

  public info(data: Omit<LogFormat, "timestamp" | "level">): void {
    this.log({ ...data, level: LogLevel.INFO });
  }

  public warn(data: Omit<LogFormat, "timestamp" | "level">): void {
    this.log({ ...data, level: LogLevel.WARN });
  }

  public error(data: Omit<LogFormat, "timestamp" | "level">): void {
    this.log({ ...data, level: LogLevel.ERROR });
  }

  public debug(data: Omit<LogFormat, "timestamp" | "level">): void {
    this.log({ ...data, level: LogLevel.DEBUG });
  }
}

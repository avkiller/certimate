export interface AccessModel extends BaseModel {
  name: string;
  provider: string;
  config?: Record<string, unknown>;
  reserve?: "ca" | "notif";
}

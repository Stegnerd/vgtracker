import toasteventbus from "primevue/toasteventbus";

// here is where this solution came from: https://stackoverflow.com/a/77720095/5745925
export function toastContrast(
  summary: string,
  detail: string,
  life: number = 3000
) {
  toasteventbus.emit("add", { severity: "contrast", summary, detail, life });
}

export function toastError(
  summary: string,
  detail: string,
  life: number = 3000
) {
  toasteventbus.emit("add", { severity: "error", summary, detail, life });
}

export function toastInfo(
  summary: string,
  detail: string,
  life: number = 3000
) {
  toasteventbus.emit("add", { severity: "info", summary, detail, life });
}

export function toastSecondary(
  summary: string,
  detail: string,
  life: number = 3000
) {
  toasteventbus.emit("add", { severity: "secondary", summary, detail, life });
}

export function toastSuccess(
  summary: string,
  detail: string,
  life: number = 3000
) {
  toasteventbus.emit("add", { severity: "success", summary, detail, life });
}

export function toastWarn(
  summary: string,
  detail: string,
  life: number = 3000
) {
  toasteventbus.emit("add", { severity: "warn", summary, detail, life });
}

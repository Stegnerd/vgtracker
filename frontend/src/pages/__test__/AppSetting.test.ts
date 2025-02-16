import { mount } from "@vue/test-utils";
import { Toast } from "primevue";
import PrimeVue from "primevue/config";
import ToastService from "primevue/toastservice";
import { vi } from "vitest";
import * as ConfigController from "../../../wailsjs/go/controllers/ConfigController"; // Adjust the import path as needed
import { config } from "../../../wailsjs/go/models";
import AppSettings from "../AppSettings.vue";
// mock the api module
vi.mock("../../../wailsjs/go/controllers/ConfigController");

describe("AppSetting", () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  test("renders properly", () => {
    vi.mocked(ConfigController.GetConfig).mockResolvedValue({
      twitch: {
        clientID: "sample-id",
        clientSecret: "sample-secret"
      }
    } as config.Config);

    const wrapper = mount(AppSettings, {
      global: {
        plugins: [PrimeVue, ToastService],
        components: {
          Toast
        }
      }
    });
    expect(wrapper.text()).toContain("Twitch");
  });
});

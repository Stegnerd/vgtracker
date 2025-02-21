import { flushPromises, mount } from "@vue/test-utils";
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
    expect(wrapper).toBeTruthy();
  });

  test("loads data", async () => {
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

    await flushPromises();

    const clientIDInput = wrapper.get('[data-test="twitchClientID"]')
      .element as HTMLInputElement;
    const clientSecretInput = wrapper.get('[data-test="twitchClientSecret"]')
      .element as HTMLInputElement;

    expect(clientIDInput.value).toBe("sample-id");
    expect(clientSecretInput.value).toBe("sample-secret");
  });
});

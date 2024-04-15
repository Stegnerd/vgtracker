import { createTestingPinia } from "@pinia/testing";
import { fireEvent, render, screen } from "@testing-library/vue";
import { fn } from "@vitest/spy";
import { setActivePinia } from "pinia";
import Button from "primevue/button";
import PrimeVue from "primevue/config";
import FloatLabel from "primevue/floatlabel";
import InputText from "primevue/inputtext";
import { Field, Form } from "vee-validate";

import { ReadConfigOutput } from "../../../src-tauri/bindings/config/ReadConfigOutput";
import { UpdateConfigInput } from "../../../src-tauri/bindings/config/UpdateConfigInput";
import { useConfigStore } from "../../stores/configStore";
import SettingView from "../../views/SettingView.vue";

describe("SettingView Tests", () => {
  const emptyMountOptions = {
    global: {
      plugins: [
        createTestingPinia({
          initialState: {
            config: {
              configuration: {
                twitchClientId: "",
                twitchClientSecret: "",
                theme: "Dark"
              } as ReadConfigOutput
            }
          },
          stubActions: true
        }),
        PrimeVue
      ],
      components: {
        PrimeInputText: InputText,
        PrimeButton: Button,
        PrimeFloatLabel: FloatLabel,
        Form,
        Field
      }
    }
  };

  const filledPinaia = createTestingPinia({
    initialState: {
      config: {
        configuration: {
          twitchClientId: "myClientID",
          twitchClientSecret: "myClientSecret",
          theme: "Dark"
        } as ReadConfigOutput
      }
    },
    stubActions: true,
    createSpy: fn
  });
  const filledMountOptions = {
    global: {
      plugins: [
        filledPinaia,
        // TODO: figure out how to set primevue globally
        PrimeVue
      ],
      components: {
        PrimeInputText: InputText,
        PrimeButton: Button,
        PrimeFloatLabel: FloatLabel,
        Form,
        Field
      }
    }
  };

  test("should load data from store", async () => {
    render(SettingView, filledMountOptions);

    const twitchClientIDInput = screen.getByTestId(
      "twitchClientIDInput"
    ) as HTMLInputElement;
    expect(twitchClientIDInput.value).toBe("myClientID");

    const twitchClientSecret = screen.getByTestId(
      "twitchClientSecretInput"
    ) as HTMLInputElement;
    expect(twitchClientSecret.value).toBe("myClientSecret");

    const submitButton = screen.getByTestId("submit-button");
    // expect button to be enabled
    expect(submitButton).toHaveProperty("disabled", false);
  });

  test("validates twitchClientId input", async () => {
    render(SettingView, emptyMountOptions);

    const twitchClientInput = screen.getByTestId("twitchClientIDInput");
    await fireEvent.update(twitchClientInput, "");
    expect(
      await screen.findByText("twitchClientId is a required field")
    ).toBeTruthy();

    const twitchSecretInput = screen.getByTestId("twitchClientSecretInput");
    await fireEvent.update(twitchSecretInput, "");
    expect(
      await screen.findByText("twitchClientSecret is a required field")
    ).toBeTruthy();

    const submitButton = screen.getByTestId("submit-button");
    // expect button to be disabled
    expect(submitButton).toHaveProperty("disabled", true);
  });

  test("should save data", async () => {
    render(SettingView, filledMountOptions);

    setActivePinia(filledPinaia);
    const store = useConfigStore(filledPinaia);
    vi.mocked(store.updateConfig).mockImplementation(() => Promise.resolve());

    await fireEvent.submit(screen.getByTestId("submit-button"));

    expect(store.updateConfig).toHaveBeenCalledWith({
      twitchClientId: "myClientID",
      twitchClientSecret: "myClientSecret"
    } as UpdateConfigInput);
  });
});

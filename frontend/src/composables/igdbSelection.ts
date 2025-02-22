import { ref } from "vue";
import { igdb } from "../../wailsjs/go/models";

const gameSelection = ref<igdb.VGTGame>();
export function useIGDBSelection() {
  const setGameSelection = (input: igdb.VGTGame) =>
    (gameSelection.value = input);
  return {
    gameSelection,
    setGameSelection
  };
}

import type { Actions } from "@sveltejs/kit";

export const actions: Actions = {
    default: async (event) => {
        event.fetch()
    }
}
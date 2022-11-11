<script lang="ts">
    export let type: string = "primary";
    export let submit: boolean = false;
    export let fullwidth: boolean = false;
    export let href: string = "";

    $: btnType = submit ? "submit" : "button"
</script>

{#if href === ""}
    <button class="button {type}" class:fullwidth={fullwidth} type={btnType}>
        <slot></slot>
    </button>
{:else}
    <a class="button {type}" class:fullwidth={fullwidth} {href}>
        <slot></slot>
    </a> 
{/if}

<style>
    .button {
        color: var(--text-contrast);
        border: none;
        border-radius: 5px;
        padding: 13px 25px;
        font: inherit;
        width: 100%;
        max-width: fit-content;
        text-decoration: none;
    }
    .button:hover {
        cursor: pointer;
    }
    .button:active {
        padding: 10px 22px;
    }

    .primary {
        background-color: var(--primary-main);
    }
    .primary:hover {
        background-color: var(--primary-light);
    }
    .primary:active {
        background-color: var(--primary-main);
        border: 3px solid var(--primary-dark);
    }

    .secondary {
        background-color: var(--secondary-main);
    }
    .secondary:hover {
        background-color: var(--secondary-light);
    }
    .secondary:active {
        background-color: var(--secondary-main);
        border: 3px solid var(--secondary-dark);
    }
</style>
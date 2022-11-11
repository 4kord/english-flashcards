<script lang="ts">
    export let name: string = "";
    export let type: string = "";
    export let fullwidth: boolean = false;
</script>

<div class:fullwidth={fullwidth}>
    <label for={name}>{name}</label>
    <input class:fullwidth={fullwidth} {name} {type} />
</div>

<style>
    input {
        display: block;
        width: 100%;
        max-width: 400px;
        padding: 12px 20px;
    }
</style>
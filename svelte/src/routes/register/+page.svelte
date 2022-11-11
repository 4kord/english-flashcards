<script lang="ts">
    import Button from "$lib/components/Button.svelte";
    import Input from "$lib/components/Input.svelte";
</script>
<main>
    <form>
        <Input name="email" type="email" fullwidth/>
        <Input name="password" type="password" fullwidth/>
        <Input name="repeat_password" type="password" fullwidth/>
        <Button submit>Register</Button>
    </form>
</main>

<style>
    main {
        padding: 0 10px;
        width: 100%;
        display: flex; 
        justify-content: center;
    }

    form {
        width: 100%;
        max-width: 500px;
        display: flex;
        flex-direction: column;
        gap: 20px;
    }
</style>
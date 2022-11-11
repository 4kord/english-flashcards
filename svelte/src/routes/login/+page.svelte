<script lang="ts">
    import Button from "$lib/components/Button.svelte";
    import Input from "$lib/components/Input.svelte";

    import { enhance } from "$app/forms";
</script>

<main>
    <form method="POST">
        <Input name="email" type="email" />
        <Input name="password" type="password" />
        <Button submit>Login</Button>
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
<script>
    import Button from "$lib/components/Button.svelte";
</script>

<main>
    <div class="content">
        <h1>English Flashcards</h1>
        <p>Lorem ipsum dolor sit amet consectetur, adipisicing elit. Ducimus impedit, soluta quaerat, quam officiis quibusdam libero vitae obcaecati id perspiciatis praesentium eaque necessitatibus aspernatur sunt reiciendis tempora! Iste, dolor similique?</p>
        <Button>Login</Button>
    </div>
</main>

<style>
    h1, p {
        text-align: left;
    }
    main {
        padding: 0 100px;
        height:700px;
        background-color: #281454;
    }
    .content {
        width: 100%;
        height: 100%;
        display: flex;
        flex-direction: column;
        justify-content: center;
        gap: 25px;
        max-width: 400px;
    }

    @media screen and (max-width: 600px) {
        main {
            padding: 0 20px;
        }
    }
</style>
<script lang="ts">
    import "./style.css"
    let anchors = [
        {
            content: "Dashboard",
            path: "/dashboard"
        },
        {
            content: "Calendar",
            path: "/calendar"
        },
        {
            content: "Payslip",
            path: "/payslip"
        },
        {
            content: "Manage",
            path: "/manage"
        }
    ]
</script>


<nav id="nav-wrapper">
    <div id="anchors-wrapper">
        {#each anchors as a}
            <a href="#{a.path}" class="anchor-nav">{a.content}</a>
        {/each}
    </div>
</nav>
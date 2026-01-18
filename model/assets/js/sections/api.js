registerSection({
  id: "api",
  title: "Query Builder API",
  content: `
    <h2>Query Builder API</h2>

    <h3>Query Initiation</h3>
    <pre><code>.Get()
.Create()
.Update(field)</code></pre>

    <h3>Conditions</h3>
    <pre><code>.Where(field).Is(value)
.Where(field).In(values...)
.Where(field).Between(a, b)</code></pre>
  `
});

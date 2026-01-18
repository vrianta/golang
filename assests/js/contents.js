const docs = [
  {
    id: "intro",
    title: "Introduction",
    content: `
      <p><strong>ModelsHandler</strong> is a human-friendly ORM-like query builder for Go.</p>
      <p>It provides chainable APIs, schema synchronization, and readable queries.</p>
    `
  },
  {
    id: "features",
    title: "Features",
    content: `
      <ul>
        <li>Fluent query builder</li>
        <li>Auto schema migration</li>
        <li>Type-safe fields</li>
        <li>MySQL support</li>
      </ul>
    `
  },
  {
    id: "install",
    title: "Installation",
    content: `
      <pre><code>go get github.com/vrianta/golang/model</code></pre>
    `
  },
  {
    id: "models",
    title: "Defining Models",
    content: `
      <pre><code>
var Users = model.New(db, "users", struct {
    UserId *model.Field
    Email  *model.Field
}{
    UserId: model.CreateField().AsBigInt().IsPrimary(),
    Email:  model.CreateField().AsVarchar(255).IsUnique(),
})
      </code></pre>
    `
  },
  {
    id: "queries",
    title: "Query Builder",
    content: `
      <pre><code>
Users.Get().
  Where(Users.Fields.Email).Like("%gmail.com").
  Limit(10).
  Fetch()
      </code></pre>
    `
  }
]

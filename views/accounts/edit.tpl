<section class="profile">
  <h1>Edit Profile</h1>
  {{if .updateError}}
      <div class="errors">
        <ul>
          <li>{{.updateError}}</li>
        <ul>
      </div>
    {{end}}
  <div class="black-block">
    <form action="/{{.accountname}}" method="POST">
      {{.xsrf}}
      <input type="hidden" name="_method" value="PUT">
      <input type="text" class="form-underbar-input" placeholder="{{.accountname}}" name="Name">
      <br>
      <input type="submit" class="form-submit" value="Change Profile!">
    </form>
    <br>
    <a href="/{{.accountname}}" class="button">Cancel</a>
    <br>
    <form action="/{{.accountname}}" method="POST">
      {{.xsrf}}
      <input type="hidden" name="_method" value="DELETE">
      <input type="submit" class="form-submit" value="Delete Account">
    </form>
  </div>
</section>
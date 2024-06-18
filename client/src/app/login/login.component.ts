import { CommonModule } from "@angular/common";
import { Component } from "@angular/core";
import { FormControl, FormGroup, ReactiveFormsModule } from "@angular/forms";
// biome-ignore lint/style/useImportType: <explanation>
import { LoginService } from "../login.service";

@Component({
	selector: "app-login",
	standalone: true,
	imports: [CommonModule, ReactiveFormsModule],
	templateUrl: "./login.component.html",
	styleUrl: "./login.component.css",
})
export class LoginComponent {
	ImagePath: string;
	applyForm = new FormGroup({
		login: new FormControl(""),
		password: new FormControl(""),
	});

	constructor(private loginService: LoginService) {
		this.ImagePath = " ../../assets/edina_logo.svg";
	}

	submitApplication() {
		this.loginService.submitApplication(
			this.applyForm.value.login ?? "",
			this.applyForm.value.password ?? "",
		);
	}
}

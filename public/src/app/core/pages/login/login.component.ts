import { NOT_FOUND, WRONG_CRED } from './../../services/auth.service';
import {
  routeName as dashboardRouteName,
  routeName,
} from './../../../dashboard/dashboard-routing.module';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../../services/auth.service';
import { PopupService } from '../../services/popup.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
})
export class LoginComponent implements OnInit {
  constructor(
    private router: Router,
    private authSerivce: AuthService,
    private popupService: PopupService
  ) {}

  ngOnInit(): void {}

  onSuccess(): void {
    this.router.navigate(['/' + routeName]);
  }
}

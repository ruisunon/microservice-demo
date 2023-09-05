package com.rydata.springboot3.controllers;

import com.rydata.springboot3.entities.Doctor;
import com.rydata.springboot3.entities.Patient;
import com.rydata.springboot3.services.DoctorPatientService;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("hospital/")
public class HospitalController {
    
    @Autowired
    DoctorPatientService doctorPatientService;

    @PostMapping("adddoctor")
    public Doctor addDoctor(@RequestBody Doctor doctor){
        return doctorPatientService.addDoctor(doctor);
    }

    @PostMapping("addpatient")
    public Patient addPatient(@RequestBody Patient patient){
        return doctorPatientService.addPatient(patient);
    }
}

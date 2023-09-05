package com.rydata.springboot3.services;

import com.rydata.springboot3.entities.Doctor;
import com.rydata.springboot3.entities.Patient;
import com.rydata.springboot3.repositories.DoctorRepository;
import com.rydata.springboot3.repositories.PatientRepository;


import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service

public class DoctorPatientService {

    @Autowired
    DoctorRepository doctorRepository;
    @Autowired
    PatientRepository patientRepository;

    public DoctorPatientService(){

    }

    public Doctor addDoctor(Doctor doctor){
        return doctorRepository.save(doctor);
    }
    
    public Patient addPatient(Patient patient){
        return patientRepository.save(patient);
    }
}

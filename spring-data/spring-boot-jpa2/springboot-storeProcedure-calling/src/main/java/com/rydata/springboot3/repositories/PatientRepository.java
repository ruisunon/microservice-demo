package com.rydata.springboot3.repositories;

import com.rydata.springboot3.entities.Patient;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface PatientRepository extends JpaRepository<Patient,Integer>{
    
}

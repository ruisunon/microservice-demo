package com.rydata.springboot3.repositories;

import com.rydata.springboot3.entities.Doctor;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface DoctorRepository extends JpaRepository<Doctor,Integer> {
    
}

package com.rydata.springboot3.repositories.sms;

import com.rydata.springboot3.entities.sms.Inbox;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface InboxRepository extends JpaRepository<Inbox,Integer> {
    
}

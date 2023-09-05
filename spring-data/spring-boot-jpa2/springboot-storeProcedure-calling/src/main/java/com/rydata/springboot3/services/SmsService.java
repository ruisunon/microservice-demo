package com.rydata.springboot3.services;

import java.util.List;

import com.rydata.springboot3.entities.Sms;
import com.rydata.springboot3.entities.sms.Failed;
import com.rydata.springboot3.entities.sms.Inbox;
import com.rydata.springboot3.entities.sms.Other;
import com.rydata.springboot3.entities.sms.Sent;
import com.rydata.springboot3.repositories.SmsRepository;
import com.rydata.springboot3.repositories.sms.FailedRepository;
import com.rydata.springboot3.repositories.sms.InboxRepository;
import com.rydata.springboot3.repositories.sms.OtherRepository;
import com.rydata.springboot3.repositories.sms.SentRepository;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class SmsService {

    @Autowired
    SmsRepository smsRepository;

    @Autowired
    InboxRepository inboxRepository;

    @Autowired
    SentRepository sentRepository;

    @Autowired
    FailedRepository failedRepository;

    @Autowired
    OtherRepository otherRepository;

    public SmsService(){


    }

    public Inbox addToInbox(Inbox inbox) {
        return inboxRepository.save(inbox);
    }

    public Sent addToSent(Sent sent) {
        return sentRepository.save(sent);
    }

    public Failed addToFailed(Failed failed) {
        return failedRepository.save(failed);
    }

    public List<Inbox> fetchInbox() {
        return inboxRepository.findAll();
    }

    public List<Sms> fetchAll() {
        return smsRepository.findAll();
    }

    public List<Other> fetchOthers() {
        return otherRepository.findAll();
    }
    
    
}

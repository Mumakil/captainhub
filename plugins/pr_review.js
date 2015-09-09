function getReviewersForFile(files, pullRequestUser, preferredReviewers) {
  var reviewerFiles = {};
  var fileReviewers = {};

  _.each(preferredReviewers, function (files, reviewer) {
    _.each(files, function (file) {
      fileReviewers[file] = fileReviewers[file] || [];
      fileReviewers[file].push(reviewer);
    });
  });

  _.each(files, function (file) {
    var reviewers = preferredReviewers[file.filename] || [];
    if (reviewers.length > 0) {
      _.each(reviewers, function (reviewer) {
        reviewerFiles[reviewer] = reviewerFiles[reviewer] || [];
        reviewerFiles[reviewer].push(reviewer);
      });
    } else {
      var pattern = _.find(config.patterns, function (fileConf) {
        var patterns = fileConf.pattern;
        var usernames = fileConf.reviewers;

        var match = _.find(patterns.split(','), function (pattern) {
          //console.log(pattern, file.filename);
          return matchFilePath(pattern, file.filename);
        });

        // Reviewer for file cant be the same as the pull request sender
        if (match && (usernames.indexOf(pullRequestUser) === -1 || usernames.length > 1)) {
          return _.without(usernames, pullRequestUser).length > 0;
        }
      });

      if (!pattern) {
        return;
      }

      reviewers = _.without(pattern.reviewers, pullRequestUser);

      var alreadyReviewing = _.intersection(
        reviewers,
        _.keys(reviewerFiles)
      );

      if (alreadyReviewing.length > 0) {
        reviewers = alreadyReviewing;
      }

      var reviewerCount = pattern.reviewerCount || 1;

      for (var i = 0; i < reviewerCount && reviewers.length > 0; i++) {
        var randomReviewer = reviewers.splice(Math.floor(Math.random() * reviewers.length), 1)[0];
        if (randomReviewer) {
          reviewerFiles[randomReviewer] = reviewerFiles[randomReviewer] || [];
          reviewerFiles[randomReviewer].push(file.filename);
        }
      }

      if (pattern.owner) {
        reviewerFiles[pattern.owner] = reviewerFiles[pattern.owner] || [];
        reviewerFiles[pattern.owner].push(file.filename);
      }
    }
  }, this);

  return reviewerFiles;
}

function getFileShaMap(files) {
  var fileShaMap = {};
  _.each(files, function (file) {
    fileShaMap[file.filename] = file.sha;
  });
  return fileShaMap;
}

function pullRequestOpened(eventData) {
  var pullRequestUser = eventData.pull_request.user.login;

  var files = getPullRequestFiles(eventData.number);
  var reviewerFiles = getReviewersForFile(files, pullRequestUser, {});
  var fileShaMap = getFileShaMap(files);

  if (_.keys(reviewerFiles).length > 0) {
    var message = 'Awesome work! Now just sit back and wait for Travis to pass and others to review your code.\n\n';
    message += '## Reviewers\n';

    _.each(reviewerFiles, function (files, reviewer) {
      message += '@' + reviewer + '\n';
      _.each(files, function (file) {
        message += '- `' + file + '`\n';
      });
      message += '\n';
    });

    message += '\n### Review commands\n';
    message += '- accept: `pr_review OK`\n';
    message += '- add reviewer: `pr_review add [username] [pattern]`\n';
    message += '- change reviewer: `pr_review change [old_username] [new_username]`\n';
    message += '- print status: `pr_review status`\n';

    createIssueComment(eventData.number, message);
    createStatus(
      eventData.pull_request.head.sha,
      'pending',
      eventData.pull_request.url,
      'Pull request review: ' + _.keys(reviewerFiles).join(', '),
      'pr_review'
    );
    saveData(eventData.number + ':waiting_review_from', JSON.stringify(reviewerFiles));
  }

  saveData(eventData.number + ':files', JSON.stringify(fileShaMap));
}

function pullRequestUpdated(eventData) {
  var pullRequestUser = eventData.pull_request.user.login;
  var files = getPullRequestFiles(eventData.number);
  var fileShaMapStr = loadData(eventData.number + ':files');
  var oldReviewersStr = loadData(eventData.number + ':waiting_review_from');
  var fileShaMap = {};
  var oldReviewers = {};

  try {
    fileShaMap = JSON.parse(fileShaMapStr);
  } catch(e) {}

  try {
    oldReviewers = JSON.parse(oldReviewersStr);
  } catch(e) {}

  files = _.filter(files, function (file) {
    if (fileShaMap[file.filename] !== file.sha) {
      fileShaMap[file.filename] = file.sha;
      return true;
    }
  });

  var reviewerFiles = getReviewersForFile(files, pullRequestUser, oldReviewers);

  if (_.keys(reviewerFiles).length > 0) {
    var message = 'Some of the files were updated.\n\n';
    message += '## Reviewers\n';
    _.each(reviewerFiles, function (files, reviewer) {
      message += '@' + reviewer + '\n';
      _.each(files, function (file) {
        message += '- `' + file + '`\n';
      });
      message += '\n';
    });

    // Merge olds files back
    _.each(oldReviewers, function (files, reviewer) {
      reviewerFiles[reviewer] = reviewerFiles[reviewer] ? _.union(files, reviewerFiles[reviewer]) : files;
    });

    createIssueComment(eventData.number, message);

    createStatus(
      eventData.pull_request.head.sha,
      'pending',
      eventData.pull_request.url,
      'Pull request review: ' + _.keys(reviewerFiles).join(', '),
      'pr_review'
    );
    saveData(eventData.number + ':waiting_review_from', JSON.stringify(reviewerFiles));
  }

  saveData(eventData.number + ':files', JSON.stringify(fileShaMap));
}

function pullRequestComment(eventData) {
  var id = eventData.issue.number;
  var senderName = eventData.sender.login;
  var body = eventData.comment.body.trim();
  var cmd;

  _.each(body.split('\n'), function (message) {
    if (message.indexOf('pr_review') === 0) {
      cmd = _.filter(
        _.map(message.substring(10).split(' '), function (part) {
          return part.trim().replace(/^@/, '');
        }),
        function (cmd) {
          return !!cmd;
        }
      );

      console.log('cmd ' + cmd.join(' '));

      switch(cmd[0].toLowerCase()) {
        case 'ok':
          var strData = loadData(id + ':waiting_review_from');
          if (strData) {
            try {
              var reviewerFiles = JSON.parse(strData);
              delete reviewerFiles[senderName];
              saveData(id + ':waiting_review_from', JSON.stringify(reviewerFiles));

              var prDetails = getPullRequestDetails(id);
              if (_.keys(reviewerFiles).length === 0) {
                createStatus(prDetails.head.sha, 'success', prDetails.url, 'Pull request review done', 'pr_review');
                createIssueComment(id, "Review completed!");
              } else {
                createStatus(
                  prDetails.head.sha,
                  'pending',
                  prDetails.url,
                  'Pull request review: ' + _.keys(reviewerFiles).join(', '),
                  'pr_review'
                );
              }

            } catch(e) {}
          }
        break;
        case 'add':
          if (cmd[1]) {
            var newFiles = [];
            var strData = loadData(id + ':waiting_review_from');
            if (strData) {
              try {
                var reviewerFiles = JSON.parse(strData);

                if (cmd[2]) {
                  newFiles = _map(_.filter(getPullRequestFiles(id), function (file) {
                    return matchFilePath(cmd[2], file.filename);
                  }), function (file) {
                    return file.filename;
                  });
                }

                reviewerFiles[cmd[1]] = _.union(reviewerFiles[cmd[1]] || [], newFiles);
                saveData(id + ':waiting_review_from', JSON.stringify(reviewerFiles));
                var prDetails = getPullRequestDetails(id);
                createStatus(
                  prDetails.head.sha,
                  'pending',
                  prDetails.url,
                  'Pull request review: ' + _.keys(reviewerFiles).join(', '),
                  'pr_review'
                );
                createIssueComment(id, '@' + cmd[1] + ' added');
              } catch(e) {}
            }
          }
        break;
        /*
        case 'remove':
          if (cmd[1]) {
            var strData = loadData(id + ':waiting_review_from');
            if (strData) {
              try {
                var reviewerFiles = JSON.parse(strData);
                delete reviewerFiles[cmd[1]];
                saveData(id + ':waiting_review_from', JSON.stringify(reviewerFiles));
              } catch(e) {}
            }
          }
        break;
        */
        case 'change':
          if (cmd[1] && cmd[2]) {
            var strData = loadData(id + ':waiting_review_from');
            if (strData) {
              try {
                var reviewerFiles = JSON.parse(strData);

                if (reviewerFiles[cmd[1]]) {
                  reviewerFiles[cmd[2]] = _.union(reviewerFiles[cmd[2]] || [], reviewerFiles[cmd[1]]);
                  delete reviewerFiles[cmd[1]];
                  saveData(id + ':waiting_review_from', JSON.stringify(reviewerFiles));
                  createIssueComment(id, '@' + cmd[1] + ' changed to @' + cmd[2]);
                }
              } catch(e) {}
            }
          }
        break;
        case 'reopen':
        var fakeEvent = {number: id, pull_request: getPullRequestDetails(id)};
        pullRequestOpened(fakeEvent);
        break;
        case 'status':
          var strData = loadData(id + ':waiting_review_from');
          if (strData) {
            try {
              var reviewerFiles = JSON.parse(strData);
              var message = 'Waiting review from\n\n';
              _.each(reviewerFiles, function (files, reviewer) {
                message += '@' + reviewer + '\n';
                _.each(files, function (file) {
                  message += '- ' + file + '\n';
                });
                message += '\n';
              });

              createIssueComment(id, message);

              if (_.keys(reviewerFiles).length === 0) {
                var prDetails = getPullRequestDetails(id);
                createStatus(prDetails.head.sha, 'success', prDetails.url, 'Pull request review done', 'pr_review');
                createIssueComment(id, "Review completed!");
              } else {
                createStatus(
                  prDetails.head.sha,
                  'pending',
                  prDetails.url,
                  'Pull request review: ' + _.keys(reviewerFiles).join(', '),
                  'pr_review'
                );
              }

            } catch(e) {}
          }
        break;
      }
    }
  });
}

// Process hook data
if (eventType === 'pull_request' && eventData.action === 'opened') {
  pullRequestOpened(eventData);
} else if (eventType === 'issue_comment' && eventData.action === 'created' && eventData.issue.pull_request) {
  pullRequestComment(eventData);
} else if (eventType === 'pull_request' && eventData.action === 'synchronize') {
  pullRequestUpdated(eventData);
}
